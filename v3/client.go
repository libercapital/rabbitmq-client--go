package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/bavalogs.git"
)

type Client interface {
	NewPublisher(queueArgs *QueueArgs, exchangeArgs *ExchangeArgs) (Publisher, error)
	NewConsumer(args ConsumerArgs) (Consumer, error)
	GetConnection() *amqp.Connection
	Close() error
	OnReconnect(func())
	DirectReplyTo(ctx context.Context, exchange, key string, timeout int, messge IncomingEventMessage) (IncomingEventMessage, error)
}

type clientImpl struct {
	credential        Credential
	reconnectionDelay int

	connection        *amqp.Connection
	closed            int32
	callbackReconnect []func()
	declare           bool

	reconnecting     chan bool
	heartbeatTimeout *int
}

// IsClosed indicate closed by developer
func (client *clientImpl) IsClosed() bool {
	return (atomic.LoadInt32(&client.closed) == 1)
}

func (client *clientImpl) OnReconnect(callback func()) {
	client.callbackReconnect = append(client.callbackReconnect, callback)
}

// Close ensure closed flag set
func (client *clientImpl) Close() error {
	bavalogs.Debug(context.Background()).Stack().Msg("closing connection")

	if client.IsClosed() {
		return amqp.ErrClosed
	}

	atomic.StoreInt32(&client.closed, 1)

	return client.connection.Close()
}

func (client *clientImpl) connect() error {
	if client.heartbeatTimeout == nil {
		defaultValue := 10
		client.heartbeatTimeout = &defaultValue
	}

	conn, err := amqp.DialConfig(client.credential.GetConnectionString(), amqp.Config{
		Heartbeat: time.Duration(*client.heartbeatTimeout) * time.Second,
	})

	if err != nil {
		return err
	}

	client.connection = conn

	go func(connParam *amqp.Connection) {
		for {
			newConn, err := client.reconnect(connParam, client.credential.GetConnectionString())

			if newConn == nil {
				errMsg := errors.New("could not reconnect to rabbitmq")

				if err != nil {
					errMsg = err
				}

				bavalogs.Fatal(context.Background(), errMsg).Send()
			}

			if err != nil {
				bavalogs.Fatal(context.Background(), err).Send()
			}

			client.connection = newConn

			for _, callback := range client.callbackReconnect {
				callback()
			}
		}
	}(conn)

	return nil
}

func (client *clientImpl) reconnect(connParam *amqp.Connection, credentials string) (*amqp.Connection, error) {
	var err error
	retries := 0

	chanErr := <-client.connection.NotifyClose(make(chan *amqp.Error))

	client.reconnecting = make(chan bool)
	defer func() {
		close(client.reconnecting)
		client.reconnecting = nil
	}()

	for {
		if retries >= 60 {
			err = fmt.Errorf("could not reconnect to rabbitmq after %d retries", retries)
			break
		}

		bavalogs.Debug(context.Background()).Interface("chan_err", chanErr).Msg("rabbitmq connection lost, trying reconnect")

		time.Sleep(time.Second)

		connParam, err := amqp.Dial(client.credential.GetConnectionString())

		if err != nil {
			bavalogs.Warn(context.Background()).Err(err).Msg("error rabbitmq trying reconnect")
			retries++
			continue
		}

		return connParam, nil
	}

	return nil, err
}

func New(credential Credential, options ClientOptions) (Client, error) {
	client := &clientImpl{
		credential:        credential,
		reconnectionDelay: options.ReconnectionDelay,
		declare:           options.Declare,
		heartbeatTimeout:  options.HeartbeatTimeout,
	}

	err := client.connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client *clientImpl) NewPublisher(queueArgs *QueueArgs, exchangeArgs *ExchangeArgs) (Publisher, error) {
	publish := publisherImpl{
		queueArgs:    queueArgs,
		exchangeArgs: exchangeArgs,
		client:       client,
		declare:      client.declare,
	}
	err := publish.connect()
	return &publish, err
}

func (client *clientImpl) NewConsumer(args ConsumerArgs) (Consumer, error) {
	consumer := consumerImpl{
		Args:    args,
		client:  client,
		declare: client.declare,
	}
	err := consumer.connect()
	return &consumer, err
}

func (client *clientImpl) GetConnection() *amqp.Connection {
	return client.connection
}

// DirectReplyTo publish an message into queue and expect an response RPC formart
// Error can be typeof models.TIMEOUT_ERROR
func (client *clientImpl) DirectReplyTo(ctx context.Context, exchange, key string, timeout int, message IncomingEventMessage) (event IncomingEventMessage, err error) {
	clientId := uuid.NewString()

	var timer time.Timer
	if timeout > 0 {
		timer = *time.NewTimer(time.Duration(timeout) * time.Second)
	}

	channel, err := client.connection.Channel()
	if err != nil {
		return
	}

	messages, err := channel.Consume("amq.rabbitmq.reply-to", clientId, true, false, false, false, nil)
	if err != nil {
		return
	}

	b, _ := json.Marshal(message)
	if err = channel.Publish(exchange, key, false, false, amqp.Publishing{
		ReplyTo:       "amq.rabbitmq.reply-to",
		CorrelationId: message.CorrelationID,
		Expiration:    strconv.Itoa(timeout * 1000),
		Body:          b,
	}); err != nil {
		return
	}

	defer func() {
		channel.Cancel(clientId, false)
		channel.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			return event, TIMEOUT_ERROR
		case msg := <-messages:
			if msg.CorrelationId == message.CorrelationID {
				var body []byte
				if msg.Body != nil {
					body = msg.Body
				}

				err = json.Unmarshal(body, &event)

				return
			}
		}
	}
}
