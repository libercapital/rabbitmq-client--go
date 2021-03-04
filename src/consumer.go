package src

import "github.com/streadway/amqp"

/**
@See https://github.com/streadway/amqp/blob/master/channel.go
	 Consume method.
*/

type ConsumerConfig struct {
	QueueName string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func Consume(config *ConsumerConfig, channel *amqp.Channel) <-chan amqp.Delivery {
	messages, err := channel.Consume(
		config.QueueName,
		config.Consumer,
		config.AutoAck,
		config.Exclusive,
		config.NoLocal,
		config.NoWait,
		config.Args,
	)

	FailOnError(err, "Failed to register a consumer")

	return messages
}
