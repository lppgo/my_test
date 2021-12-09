package stomp

import (
	"go-micro.dev/v4/broker"
	"github.com/go-stomp/stomp/v3"
)

type subscriber struct {
	opts  broker.SubscribeOptions
	topic string
	sub   *stomp.Subscription
}

func (s *subscriber) Options() broker.SubscribeOptions {
	return s.opts
}

func (s *subscriber) Topic() string {
	return s.topic
}

func (s *subscriber) Unsubscribe() error {
	return s.sub.Unsubscribe()
}
