package systems

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
	Inbox    chan base.Message
	Outbox   chan base.Message
}

//Create a new InputSystem - static method
func NewInputSystem(in chan base.Message, out chan base.Message) *InputSystem {
	ThisReader := bufio.NewReader(os.Stdin)
	Input := InputSystem{Input: ThisReader, Messages: make([]base.Message, 0), Inbox: in, Outbox: out}
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

	e.RecieveMessage()

	text := e.Getline()
	e.SendMessage(LogicID, text, 0) //Send Message here
	return true
}

//Getline grabs a line from stdin
func (e *InputSystem) Getline() string {
	text, _ := e.Input.ReadString('\n')
	return text
}

//SendMessage sends messages to other systems
//Make this a channel or smt later
func (e *InputSystem) SendMessage(dest int, str string, c int) bool {
	msg := base.Message{From: InputID, To: dest, Content: str, Code: c}
	e.Outbox <- msg
	return true
}

//Process recieved messages from other systems
//Currently not in use since InputSystem doesn't respond to messages
func (e *InputSystem) RecieveMessage() bool {
	return true
}
