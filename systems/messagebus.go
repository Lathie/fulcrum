package systems

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
	"strconv"
)

//Constants to refer to Systems by ID
const (
	InputID = iota
	LogicID
)

//Inbox - Channel containing incoming messages from other systems
//Input - Channel that is the inbox for InputSystem
//Logic - Channel that is the inbox for LogicSystem
type MessageBus struct {
	Inbox chan base.Message
	Input chan base.Message
	Logic chan base.Message

	Debug bool
}

//Create a new MessageBus - Static Method
func NewMessageBus(in chan base.Message, i chan base.Message, l chan base.Message, debug bool) *MessageBus {
	thisMessageBus := MessageBus{Inbox: in, Input: i, Logic: l, Debug: debug}
	logging.Log("MessageBus", "Message Bus Initialized")
	return &thisMessageBus

}

//Update every iteration of the thread's loop
func (m *MessageBus) Update() bool {
	return m.RecieveMessage()
}

//Send a message msg to a specified channel and log it
func (m *MessageBus) SendMessage(channel chan base.Message, msg base.Message) bool {
	channel <- msg
	if m.Debug {
		infostr := "Content: " + msg.Content[:len(msg.Content)-1] + " | from " + strconv.Itoa(msg.From) + " to " + strconv.Itoa(msg.To)
		fmt.Printf("(MB) %s\n", infostr)
		logging.Log("MessageBus", infostr)
	}
	return true
}

//Recieve a message, then send it to the proper destination
func (m *MessageBus) RecieveMessage() bool {

	msg, ok := <-m.Inbox
	if ok {
		switch msg.To {
		case InputID:
			m.SendMessage(m.Input, msg)
		case LogicID:
			m.SendMessage(m.Logic, msg)
		default:
			logging.Log("MessageBus", "(ERROR) Message Bus could not deliver a message!")
		}
	} else {
		logging.Log("MessageBus", "(ERROR) Message Bus channel inbox channel closed")
		return true
	}
	return true
}
