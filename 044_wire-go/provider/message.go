package provider

type Message struct {
	msg string
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}
