package systems

import (
	//"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
	"github.com/Lathie/fulcrum/pivots"
)

//Logic systems has these fields
//Messages - a slice to hold messages from other systems
type LogicSystem struct {
	Inbox  chan base.Message
	Outbox chan base.Message
}

//NewLogicSystem creates a new LogicSystem
func NewLogicSystem(in chan base.Message, out chan base.Message) *LogicSystem {
	logic := LogicSystem{Inbox: in, Outbox: out}
	logging.Log("LogicSystem", "Logic System Initialized")

	return &logic
}

//Update() gets called each game loop
func (l *LogicSystem) Update() bool {
	l.RecieveMessage()
	return true
}

//SendMessage sends a message to another system
//Currently does nothing
func (l *LogicSystem) SendMessage(dest int, str string, c int) bool {
	msg := base.Message{From: LogicID, To: dest, Content: str, Code: c, Args: nil}
	l.Outbox <- msg
	return true
}

//RecieveMessage processes messages sent from other systems
func (l *LogicSystem) RecieveMessage() bool {

	msg, ok := <-l.Inbox
	if ok {
		l.ParseMessage(msg)
	} else {
		logging.Log("LogicSystem", "Inbox Channel closed!")
	}

	return true
}

func (l *LogicSystem) ParseMessage(msg base.Message) bool {
	if msg.From == InputID {
		switch msg.Content {
		//Consider dynamic hotkeys?
		case "north", "n":
			l.SendMessage(WorldID, "Move North", pivots.North)
		case "east", "e":
			l.SendMessage(WorldID, "Move East", pivots.East)
		case "south", "s":
			l.SendMessage(WorldID, "Move South", pivots.South)
		case "west", "w":
			l.SendMessage(WorldID, "Move West", pivots.West)
		case "up", "u":
			l.SendMessage(WorldID, "Move Up", pivots.Up)
		case "down", "d":
			l.SendMessage(WorldID, "Move Down", pivots.Down)
		case "look", "l":
			l.SendMessage(WorldID, "Look", pivots.Look)
		case "look all", "la":
			l.SendMessage(WorldID, "Look All", pivots.LookAll)
		case "save":
			l.SendMessage(WorldID, "Save Map", pivots.SaveWorld)
		}
		return true
	} else {
		logging.Log("LogicSystem", "LogicSystem encountered a message not from InputSystem")
	}
	logging.Log("LogicSystem", "Command not recognized")

	//Here tell output to tell the user that they fucked up
	return true
}
