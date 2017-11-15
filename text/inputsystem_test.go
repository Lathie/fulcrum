package text_test

import (
	"github.com/Lathie/fulcrum/text"
	"testing"
)

//Test that the method to create new InputSystems works correctly
func TestNewInputSystem(t *testing.T) {
	logic := text.NewLogicSystem()
	input := text.NewInputSystem(logic)
	assert(t, input != nil, "InputSystem creation failed")
}

//Test that the Update() function for input systems works correctly
func TestInputSystemUpdate(t *testing.T) {
	logic := text.NewLogicSystem()
	input := text.NewInputSystem(logic)
	input.Update()
	assert(t, logic.Messages[0].Code == 0, "Message not sent to LogicSystem")
}

//Test that the GetLine() function for input systems works correctly
func TestGetLine(t *testing.T) {
	logic := text.NewLogicSystem()
	input := text.NewInputSystem(logic)
	line := input.Getline()
	assert(t, line == "", "Getline did not return as expected")
}

//Test that the SendMessage() function for input systems works correctly
func TestSendMessageInput(t *testing.T) {
	logic := text.NewLogicSystem()
	input := text.NewInputSystem(logic)
	input.SendMessage(logic, "test")
	assert(t, logic.Messages[0].Message == "test", "Message was not sent to LogicSystem")
}

//Test that the RecieveMessage() function works correctly
func TestRecieveMessageInput(t *testing.T) {
	logic := text.NewLogicSystem()
	input := text.NewInputSystem(logic)
	assert(t, input.RecieveMessage(), "RecieveMessage() did not succeed")
}
