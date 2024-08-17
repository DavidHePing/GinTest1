//go:build wireinject

package setup

import (
	"GinTest1/log"
	"fmt"

	"github.com/google/wire"
)

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Greeter struct {
	Message Message // <- adding a Message field
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
	fileLogger := log.InitLogger()
	fileLogger.Info(string(msg))
}

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
