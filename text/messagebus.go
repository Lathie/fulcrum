package text

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
)

const (
	InputID = iota
	LogicID
)

type MessageBus struct {
	Inbox chan base.Message
	Input chan base.Message
	Logic chan base.Message

	Debug bool
}

func NewMessageBus(in chan base.Message, i chan base.Message, l chan base.Message, debug bool) *MessageBus {
	thisMessageBus := MessageBus{Inbox: in, Input: i, Logic: l, Debug: debug}

	return &thisMessageBus

}

func (m *MessageBus) Update() bool {
	return m.RecieveMessage()
}

func (m *MessageBus) SendMessage(channel chan base.Message, msg base.Message) bool {
	channel <- msg
	if m.Debug {
		fmt.Printf("(MB) Content: %s", msg.Content)
	}
	return true
}

func (m *MessageBus) RecieveMessage() bool {
	select {
	case msg, ok := <-m.Inbox:
		if ok {
			switch msg.To {
			case InputID:
				m.SendMessage(m.Input, msg)
			case LogicID:
				m.SendMessage(m.Logic, msg)
			default:
				fmt.Printf("(MB) (ERROR) Message Bus could not deliver a message!\n")
				fmt.Printf("....From: %d To: %d / Message: %s\n", msg.From, msg.To, msg.Content)
			}
		} else {
			fmt.Println("(MB) Inbox Channel closed!")
			return true
		}
	default:
		fmt.Printf("(MB) Nothing read from inbox\n")
		return true
	}
	return true
}
