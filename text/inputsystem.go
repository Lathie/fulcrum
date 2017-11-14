package text

import (
	"bufio"
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"os"
)

type InputSystem struct {
	Input    *bufio.Reader
	Messages []string
}

func (e *InputSystem) Init() {
	e.Input = bufio.NewReader(os.Stdin)
}

func (e *InputSystem) Update() {
	for len(e.Messages) > 0 {
		e.RecieveMessage()
	}

	text = e.Getline()
	e.SendMessage(LogicSystem, text)
}

func (e *InputSystem) Getline() string {
	text, _ = e.Input.ReadString('\n')
	return text
}

//Make this a channel or smt later
func (e *InputSystem) SendMessage(s *System, str string) {
	s.Messages = append(s.Messages, Message{Sys: "input", Message: str})
}

func (e *InputSystem) RecieveMessage() {

}
