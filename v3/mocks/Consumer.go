// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	amqp091 "github.com/rabbitmq/amqp091-go"

	mock "github.com/stretchr/testify/mock"

	rabbitmq "github.com/libercapital/rabbitmq-client-go/v3"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// GetArgs provides a mock function with given fields:
func (_m *Consumer) GetArgs() rabbitmq.ConsumerArgs {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetArgs")
	}

	var r0 rabbitmq.ConsumerArgs
	if rf, ok := ret.Get(0).(func() rabbitmq.ConsumerArgs); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rabbitmq.ConsumerArgs)
	}

	return r0
}

// GetQueue provides a mock function with given fields:
func (_m *Consumer) GetQueue() amqp091.Queue {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetQueue")
	}

	var r0 amqp091.Queue
	if rf, ok := ret.Get(0).(func() amqp091.Queue); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(amqp091.Queue)
	}

	return r0
}

// SubscribeEvents provides a mock function with given fields: ctx, consumerEvent
func (_m *Consumer) SubscribeEvents(ctx context.Context, consumerEvent rabbitmq.ConsumerEvent) error {
	ret := _m.Called(ctx, consumerEvent)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeEvents")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, rabbitmq.ConsumerEvent) error); ok {
		r0 = rf(ctx, consumerEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeEventsWithHealthCheck provides a mock function with given fields: ctx, consumerEvent, concurrency, publisher
func (_m *Consumer) SubscribeEventsWithHealthCheck(ctx context.Context, consumerEvent rabbitmq.ConsumerEvent, concurrency int, publisher rabbitmq.Publisher) error {
	ret := _m.Called(ctx, consumerEvent, concurrency, publisher)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeEventsWithHealthCheck")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, rabbitmq.ConsumerEvent, int, rabbitmq.Publisher) error); ok {
		r0 = rf(ctx, consumerEvent, concurrency, publisher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewConsumer creates a new instance of Consumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}