package text

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
)

//Logic systems has these fields
//Messages - a slice to hold messages from other systems
type LogicSystem struct {
	Messages []base.Message
}

//NewLogicSystem creates a new LogicSystem
func NewLogicSystem() *LogicSystem {
	Logic := LogicSystem{Messages: make([]base.Message, 0)}
	return &Logic
}

//Update() gets called each game loop
func (l *LogicSystem) Update() bool {
	cont := true
	for len(l.Messages) > 0 {
		l.RecieveMessage()
	}
	return cont
}

//SendMessage sends a message to another system
func (l *LogicSystem) SendMessage(s *base.System, str string) bool {
	return true
}

//RecieveMessage processes messages sent from other systems
func (l *LogicSystem) RecieveMessage() bool {
	cont := true
	//Get initial message
	if len(l.Messages) > 0 {
		msg := l.Messages[0]
		l.Messages = l.Messages[1:]
		if msg.Code == 0 {
			fmt.Printf("InputSystem says: %s\n", msg.Message)
			if msg.Message == "quit\n" {
				cont = false
			}

		}
	}
	return cont
}
