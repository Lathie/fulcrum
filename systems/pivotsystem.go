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
		case pivots.North, pivots.East, pivots.South, pivots.West, pivots.Up, pivots.Down:
			p.Move(msg.Code)
		case pivots.Look:
			fmt.Println(p.Hero.Location)
			fmt.Printf("%s | %s \n", p.CurrentRoom().Name, p.CurrentRoom().Desc)
		case pivots.LookAll:
			fmt.Println(p.Hero.Location)
			fmt.Printf("%s | %s \n%s\n", p.CurrentRoom().Name, p.CurrentRoom().Desc, p.CurrentRoom().DescAll)

		}
	}
	return true
}

func (p *PivotSystem) CurrentRoom() pivots.Room {
	room := p.Realm.Map[p.Hero.Location[0]][p.Hero.Location[1]][p.Hero.Location[2]]
	return room
}

func (p *PivotSystem) Move(direction int) bool {
	if p.CurrentRoom().Exits[direction] {
		switch direction {
		case pivots.North:
			p.Hero.Location[0] += 1
		case pivots.East:
			p.Hero.Location[1] += 1
		case pivots.South:
			p.Hero.Location[0] -= 1
		case pivots.West:
			p.Hero.Location[1] -= 1
		case pivots.Up:
			p.Hero.Location[2] += 1
		case pivots.Down:
			p.Hero.Location[2] -= 1
		}
	} else {
		fmt.Println("No exit in that direction")
	}
	return true
}
