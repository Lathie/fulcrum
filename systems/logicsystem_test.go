package systems_test

import (
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/systems"
	"testing"
)

//Test that the method to create new LogicSystems works correctly
func TestNewLogicSystem(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)

	logic := systems.NewLogicSystem(in, out)
	assert(t, logic != nil, "LogicSystem creation failed")

	close(in)
	close(out)
}

//Test that the Update() function for LogicSystem works correctly
func TestLogicSystemUpdate(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)
	logic := systems.NewLogicSystem(in, out)

	msg := base.Message{From: 0, To: 1, Content: "Testing", Code: 0}
	in <- msg

	assert(t, logic.Update(), "LogicSystem Update() method failed")

	close(in)
	close(out)
}

//Test that the SendMessage() function for LogicSystem works correctly
func TestSendMessageLogic(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)

	logic := systems.NewLogicSystem(in, out)
	assert(t, logic.SendMessage(), "LogicSystem SendMessage() failed")

	close(in)
	close(out)
}

//Test that the RecieveMessage() function for LogicSystem works correctly
func TestRecieveMessageLogic(t *testing.T) {
	in := make(chan base.Message, 10)
	out := make(chan base.Message, 10)

	msg := base.Message{From: 0, To: 1, Content: "Testing", Code: 0}
	in <- msg

	logic := systems.NewLogicSystem(in, out)
	assert(t, logic.RecieveMessage(), "LogicSystem SendMessage() failed")

	close(in)
	close(out)
}
