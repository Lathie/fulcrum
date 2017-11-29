package systems

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
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

	logging.Log("LogicSystem", "Logic System Initialized")

	return &Logic
}

//Update() gets called each game loop
func (l *LogicSystem) Update() bool {
	l.RecieveMessage()
	return true
}

//SendMessage sends a message to another system
//Currently does nothing
func (l *LogicSystem) SendMessage() bool {
	return true
}

//RecieveMessage processes messages sent from other systems
func (l *LogicSystem) RecieveMessage() bool {

	msg, ok := <-l.Inbox
	if ok {
		fmt.Printf("(LS) IS to LS: %s", msg.Content)
		logging.Log("LogicSystem", "Message Recieved")
	} else {
		fmt.Println("(LS) Inbox Channel closed!")
	}

	return true
}
