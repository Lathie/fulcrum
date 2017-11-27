package text

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
)

//Logic systems has these fields
//Messages - a slice to hold messages from other systems
type LogicSystem struct {
	Messages []base.Message
	Inbox    chan base.Message
	Outbox   chan base.Message
}

//NewLogicSystem creates a new LogicSystem
func NewLogicSystem(in chan base.Message, out chan base.Message) *LogicSystem {
	Logic := LogicSystem{Messages: make([]base.Message, 0), Inbox: in, Outbox: out}
	return &Logic
}

//Update() gets called each game loop
func (l *LogicSystem) Update() bool {

	l.RecieveMessage()
	return true
}

//SendMessage sends a message to another system
func (l *LogicSystem) SendMessage() bool {
	return true
}

//RecieveMessage processes messages sent from other systems
func (l *LogicSystem) RecieveMessage() bool {

	select {
	case msg, ok := <-l.Inbox:
		if ok {
			fmt.Printf("(LS) IS to LS: %s", msg.Content)
		} else {
			fmt.Println("(LS) Inbox Channel closed!")
		}
	default:
		return true
	}
	return true
}
