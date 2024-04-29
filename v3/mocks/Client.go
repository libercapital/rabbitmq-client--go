// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	amqp091 "github.com/rabbitmq/amqp091-go"

	mock "github.com/stretchr/testify/mock"

	rabbitmq "github.com/libercapital/rabbitmq-client-go/v3"

	tracing "github.com/libercapital/liber-logger-go/tracing"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Client) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DirectReplyTo provides a mock function with given fields: ctx, exchange, key, timeout, messge, trace
func (_m *Client) DirectReplyTo(ctx context.Context, exchange string, key string, timeout int, messge rabbitmq.IncomingEventMessage, trace tracing.SpanConfig) (rabbitmq.IncomingEventMessage, error) {
	ret := _m.Called(ctx, exchange, key, timeout, messge, trace)

	if len(ret) == 0 {
		panic("no return value specified for DirectReplyTo")
	}

	var r0 rabbitmq.IncomingEventMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, rabbitmq.IncomingEventMessage, tracing.SpanConfig) (rabbitmq.IncomingEventMessage, error)); ok {
		return rf(ctx, exchange, key, timeout, messge, trace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, rabbitmq.IncomingEventMessage, tracing.SpanConfig) rabbitmq.IncomingEventMessage); ok {
		r0 = rf(ctx, exchange, key, timeout, messge, trace)
	} else {
		r0 = ret.Get(0).(rabbitmq.IncomingEventMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, rabbitmq.IncomingEventMessage, tracing.SpanConfig) error); ok {
		r1 = rf(ctx, exchange, key, timeout, messge, trace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields:
func (_m *Client) GetConnection() *amqp091.Connection {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 *amqp091.Connection
	if rf, ok := ret.Get(0).(func() *amqp091.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp091.Connection)
		}
	}

	return r0
}

// HealthCheck provides a mock function with given fields: publisher
func (_m *Client) HealthCheck(publisher rabbitmq.Publisher) bool {
	ret := _m.Called(publisher)

	if len(ret) == 0 {
		panic("no return value specified for HealthCheck")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(rabbitmq.Publisher) bool); ok {
		r0 = rf(publisher)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewConsumer provides a mock function with given fields: args
func (_m *Client) NewConsumer(args rabbitmq.ConsumerArgs) (rabbitmq.Consumer, error) {
	ret := _m.Called(args)

	if len(ret) == 0 {
		panic("no return value specified for NewConsumer")
	}

	var r0 rabbitmq.Consumer
	var r1 error
	if rf, ok := ret.Get(0).(func(rabbitmq.ConsumerArgs) (rabbitmq.Consumer, error)); ok {
		return rf(args)
	}
	if rf, ok := ret.Get(0).(func(rabbitmq.ConsumerArgs) rabbitmq.Consumer); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rabbitmq.Consumer)
		}
	}

	if rf, ok := ret.Get(1).(func(rabbitmq.ConsumerArgs) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPublisher provides a mock function with given fields: queueArgs, exchangeArgs
func (_m *Client) NewPublisher(queueArgs *rabbitmq.QueueArgs, exchangeArgs *rabbitmq.ExchangeArgs) (rabbitmq.Publisher, error) {
	ret := _m.Called(queueArgs, exchangeArgs)

	if len(ret) == 0 {
		panic("no return value specified for NewPublisher")
	}

	var r0 rabbitmq.Publisher
	var r1 error
	if rf, ok := ret.Get(0).(func(*rabbitmq.QueueArgs, *rabbitmq.ExchangeArgs) (rabbitmq.Publisher, error)); ok {
		return rf(queueArgs, exchangeArgs)
	}
	if rf, ok := ret.Get(0).(func(*rabbitmq.QueueArgs, *rabbitmq.ExchangeArgs) rabbitmq.Publisher); ok {
		r0 = rf(queueArgs, exchangeArgs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rabbitmq.Publisher)
		}
	}

	if rf, ok := ret.Get(1).(func(*rabbitmq.QueueArgs, *rabbitmq.ExchangeArgs) error); ok {
		r1 = rf(queueArgs, exchangeArgs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnReconnect provides a mock function with given fields: _a0
func (_m *Client) OnReconnect(_a0 func()) {
	_m.Called(_a0)
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
