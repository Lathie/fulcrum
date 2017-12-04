package systems

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
	"github.com/Lathie/fulcrum/pivots"
)

type PivotSystem struct {
	Inbox  chan base.Message
	Outbox chan base.Message

	Hero  *pivots.Player
	Realm *pivots.World
}

func NewPivotSystem(in chan base.Message, out chan base.Message) *PivotSystem {
	worldMap := pivots.NewWorld()
	player := pivots.NewPlayer(true)
	pivot := PivotSystem{Inbox: in, Outbox: out, Realm: &worldMap, Hero: &player}
	logging.Log("PivotSystem", "Pivot System Initialized")
	return &pivot
}

func (p *PivotSystem) Update() bool {
	p.RecieveMessage()
	return true
}

func (p *PivotSystem) SendMessage() bool {
	return true
}

func (p *PivotSystem) RecieveMessage() bool {
	msg, ok := <-p.Inbox
	if ok {
		p.ParseMessage(msg)
	} else {
		logging.Log("PivotSystem", "Inbox Channel closed!")
	}
	return true
}

func (p *PivotSystem) ParseMessage(msg base.Message) bool {
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
		case pivots.Look:
			fmt.Println(p.Hero.Location)
			fmt.Printf("%s | %s \n", p.CurrentRoom().Name, p.CurrentRoom().Desc)
		}
	}
	return true
}

func (p *PivotSystem) CurrentRoom() pivots.Room {
	room := p.Realm.Map[p.Hero.Location[0]][p.Hero.Location[1]][p.Hero.Location[2]]
	return room
}
