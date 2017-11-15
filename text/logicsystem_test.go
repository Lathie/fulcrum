package text_test

import (
	"github.com/Lathie/fulcrum/text"
	"testing"
)

//Test that the method to create new LogicSystems works correctly
func TestNewLogicSystem(t *testing.T) {
	logic := text.NewLogicSystem()
	assert(t, logic != nil, "LogicSystem creation failed")
}

//Test that the Update() function for LogicSystem works correctly
func TestLogicSystemUpdate(t *testing.T) {
	logic := text.NewLogicSystem()
	assert(t, logic.Update(), "LogicSystem Update() method failed")
}

//Test that the SendMessage() function for LogicSystem works correctly
func TestSendMessageLogic(t *testing.T) {
	logic := text.NewLogicSystem()
	assert(t, logic.SendMessage(nil, ""), "LogicSystem SendMessage() failed")
}

//Test that the RecieveMessage() function for LogicSystem works correctly
func TestRecieveMessageLogic(t *testing.T) {
	logic := text.NewLogicSystem()
	assert(t, logic.RecieveMessage(), "LogicSystem SendMessage() failed")
}
