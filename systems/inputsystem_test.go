package systems_test

import (
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/systems"
	"testing"
)

//Test that the method to create new InputSystems works correctly
func TestNewInputSystem(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)
	input := systems.NewInputSystem(in, out)

	assert(t, input != nil, "InputSystem creation failed")

	close(in)
	close(out)
}

//Test that the Update() function for input systems works correctly
func TestInputSystemUpdate(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)
	input := systems.NewInputSystem(in, out)

	test := input.Update()

	assert(t, test, "Message not sent to LogicSystem")

	close(in)
	close(out)
}

//Test that the GetLine() function for input systems works correctly
func TestGetLine(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)
	input := systems.NewInputSystem(in, out)

	line := input.Getline()

	assert(t, line == "", "Getline did not return as expected")

	close(in)
	close(out)
}

//Test that the SendMessage() function for input systems works correctly
func TestSendMessageInput(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)
	input := systems.NewInputSystem(in, out)

	test := input.SendMessage(1, "test", 0)
	assert(t, test, "Message was not sent to LogicSystem")

	close(in)
	close(out)
}

//Test that the RecieveMessage() function works correctly
func TestRecieveMessageInput(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)
	input := systems.NewInputSystem(in, out)

	assert(t, input.RecieveMessage(), "RecieveMessage() did not succeed")

	close(in)
	close(out)
}
