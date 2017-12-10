package systems

import (
	//	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
	"strconv"
)

//Constants to refer to Systems by ID
const (
	InputID = iota
	LogicID
	WorldID
	OutputID
)

//Inbox - Channel containing incoming messages from other systems
//Input - Channel that is the inbox for InputSystem
//Logic - Channel that is the inbox for LogicSystem
type MessageBus struct {
	Inbox  chan base.Message
	Input  chan base.Message
	Logic  chan base.Message
	World  chan base.Message
	Output chan base.Message

	Debug bool
}

//Create a new MessageBus - Static Method
func NewMessageBus(in chan base.Message, i chan base.Message, l chan base.Message, w chan base.Message, o chan base.Message, debug bool) *MessageBus {
	thisMessageBus := MessageBus{Inbox: in, Input: i, Logic: l, World: w, Output: o, Debug: debug}
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
		infostr := "Content: " + msg.Content + " | from " + strconv.Itoa(msg.From) + " to " + strconv.Itoa(msg.To)
		//fmt.Printf("(MB) %s\n", infostr)
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
		case WorldID:
			m.SendMessage(m.World, msg)
		case OutputID:
			m.SendMessage(m.Output, msg)
		default:
			logging.Log("MessageBus", "(ERROR) Message Bus could not deliver a message!")
		}
	} else {
		logging.Log("MessageBus", "(ERROR) Message Bus channel inbox channel closed")
		return true
	}
	return true
}
