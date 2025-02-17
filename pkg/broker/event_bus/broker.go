package event_bus

import (
	"context"
	"errors"

	evbus "github.com/asaskevich/EventBus"
	"github.com/codfrm/cago/pkg/broker/broker"
)

type eventBusBroker struct {
	bus evbus.Bus
}

func NewEvBusBroker() broker.Broker {
	ret := &eventBusBroker{
		bus: evbus.New(),
	}
	return ret
}

func (e *eventBusBroker) Publish(ctx context.Context, topic string, data *broker.Message, opt ...broker.PublishOption) error {
	e.bus.Publish(topic, data)
	return nil
}

func (e *eventBusBroker) Subscribe(ctx context.Context, topic string, h broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	options := broker.NewSubscribeOptions(opts...)
	if options.Group != "" {
		return nil, errors.New("event_bus not support group")
	}
	return newSubscriber(e, topic, h)
}

func (e *eventBusBroker) Close() error {
	return nil
}

func (e *eventBusBroker) String() string {
	return "event_bus"
}
