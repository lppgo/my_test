package provider

import "fmt"

type Event struct {
	Greeter Greeter
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println("start : ", msg.msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}
