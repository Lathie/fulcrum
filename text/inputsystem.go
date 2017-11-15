package text

import (
	"bufio"
	//	"fmt"
	"github.com/Lathie/fulcrum/base"
	"os"
)

//InputSystem has these fields:
//Input - a pointer to a bufio.Reader
//Messages - A slice of messages to process
//Logic - A pointer to the engine's LogicSystem
type InputSystem struct {
	Input    *bufio.Reader
	Messages []base.Message
	Logic    *LogicSystem
}

//Create a new InputSystem - static method
func NewInputSystem(l *LogicSystem) *InputSystem {
	ThisReader := bufio.NewReader(os.Stdin)
	Input := InputSystem{Input: ThisReader, Messages: make([]base.Message, 0), Logic: l}
	return &Input
}

//Create a new InputSystem
//Deprecated - Phase out of the interface definition before touching this
func (e *InputSystem) Init() {
	e.Input = bufio.NewReader(os.Stdin)

}

//Update() method that gets called each game loop
//Getline then sends the message to the logicsystem
func (e *InputSystem) Update() bool {
	cont := true
	for len(e.Messages) > 0 {
		e.RecieveMessage()
	}

	text := e.Getline()
	e.SendMessage(e.Logic, text)
	return cont
}

//Getline grabs a line from stdin
func (e *InputSystem) Getline() string {
	text, _ := e.Input.ReadString('\n')
	return text
}

//SendMessage sends messages to other systems
//Make this a channel or smt later
func (e *InputSystem) SendMessage(s *LogicSystem, str string) bool {
	s.Messages = append(s.Messages, base.Message{Sys: "input", Message: str, Code: 0})
	return true
}

func (e *InputSystem) RecieveMessage() bool {
	return true
}
