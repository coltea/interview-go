package wire

type Message string

func NewMessage(s string) Message {
	return Message(s)
}

type Greeter struct {
	Message Message
}

func NewGreeter(message Message) *Greeter {
	return &Greeter{
		Message: NewMessage("dfwedwdw"),
	}
}
