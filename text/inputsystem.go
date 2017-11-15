package text

import (
	"bufio"
	//	"fmt"
	"github.com/Lathie/fulcrum/base"
	"os"
)

type InputSystem struct {
	Input    *bufio.Reader
	Messages []base.Message
	Logic    *LogicSystem
}

func NewInputSystem(l *LogicSystem) *InputSystem {
	ThisReader := bufio.NewReader(os.Stdin)
	Input := InputSystem{Input: ThisReader, Messages: make([]base.Message, 0), Logic: l}
	return &Input
}

func (e *InputSystem) Init() {
	e.Input = bufio.NewReader(os.Stdin)

}

func (e *InputSystem) Update() bool {
	cont := true
	for len(e.Messages) > 0 {
		e.RecieveMessage()
	}

	text := e.Getline()
	e.SendMessage(e.Logic, text)
	return cont
}

func (e *InputSystem) Getline() string {
	text, _ := e.Input.ReadString('\n')
	return text
}

//Make this a channel or smt later
func (e *InputSystem) SendMessage(s *LogicSystem, str string) bool {
	s.Messages = append(s.Messages, base.Message{Sys: "input", Message: str, Code: 0})
	return true
}

func (e *InputSystem) RecieveMessage() bool {
	return true
}
