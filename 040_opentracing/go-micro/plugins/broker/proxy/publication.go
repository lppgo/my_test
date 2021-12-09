package proxy

import (
	"go-micro.dev/v4/broker"
)

type publication struct {
	topic   string
	message *broker.Message
	err     error
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.message
}

func (p *publication) Ack() error {
	return nil
}

func (p *publication) Error() error {
	return p.err
}
