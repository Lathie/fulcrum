package text

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
)

type LogicSystem struct {
	Messages []base.Message
}

func (l *LogicSystem) Update() bool {
	cont := true
	for len(l.Messages) > 0 {
		l.RecieveMessage()
	}
	return cont
}

func (l *LogicSystem) SendMessage(s *base.System, str string) bool {
	return true
}

func (l *LogicSystem) RecieveMessage() bool {
	cont := true
	//Get initial message
	msg := l.Messages[0]
	l.Messages = l.Messages[1:]
	if msg.Code == 0 {
		fmt.Printf("InputSystem says: %s\n", msg.Message)
		if msg.Message == "quit\n" {
			cont = false
		}
	}
	return cont
}
