package provider

type Greeter struct {
	Message Message
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}
