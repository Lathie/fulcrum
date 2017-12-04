package systems

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
	"github.com/Lathie/fulcrum/pivots"
)

type WorldSystem struct {
	Inbox  chan base.Message
	Outbox chan base.Message

	Hero  *pivots.Player
	Realm *pivots.World
}

func NewWorldSystem(in chan base.Message, out chan base.Message) *WorldSystem {
	worldMap := pivots.NewWorld()
	player := pivots.NewPlayer(true)
	world := WorldSystem{Inbox: in, Outbox: out, Realm: &worldMap, Hero: &player}
	logging.Log("WorldSystem", "World System Initialized")
	return &world
}

func (w *WorldSystem) Update() bool {
	w.RecieveMessage()
	return true
}

func (w *WorldSystem) SendMessage() bool {
	return true
}

func (w *WorldSystem) RecieveMessage() bool {
	msg, ok := <-w.Inbox
	if ok {
		w.ParseMessage(msg)
	} else {
		logging.Log("WorldSystem", "Inbox Channel closed!")
	}
	return true
}

func (w *WorldSystem) ParseMessage(msg base.Message) bool {
	if msg.From == LogicID {
		switch msg.Code {
		case pivots.North:
			fmt.Println("North recieved by WS")
		case pivots.East:
			fmt.Println("East recieved by WS")
		case pivots.South:
			fmt.Println("South recieved by WS")
		case pivots.West:
			fmt.Println("West recieved by WS")
		case pivots.Up:
			fmt.Println("Up recieved by WS")
		case pivots.Down:
			fmt.Println("Down recieved by WS")
		case Look:
			fmt.Println(w.Hero.Location)
			fmt.Printf("%s | %s \n", w.CurrentRoom().Name, w.CurrentRoom().Desc)
		}
	}
	return true
}

func (w *WorldSystem) CurrentRoom() pivots.Room {
	room := w.Realm.Map[w.Hero.Location[0]][w.Hero.Location[1]][w.Hero.Location[2]]
	return room
}
