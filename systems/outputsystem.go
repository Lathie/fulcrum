package systems

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
)

type OutputSystem struct {
	Inbox  chan base.Message
	Outbox chan base.Message
}

func NewOutputSystem(in chan base.Message, out chan base.Message) *OutputSystem {
	output := OutputSystem{Inbox: in, Outbox: out}

	logging.Log("OutputSystem", "Output System Initialized")
	return &output
}

func (o *OutputSystem) Update() bool {
	o.RecieveMessage()
	return true
}

func (o *OutputSystem) SendMessage() bool {
	return true
}

func (o *OutputSystem) RecieveMessage() bool {
	msg, ok := <-o.Inbox
	if ok {
		o.ParseMessage(msg)
	} else {
		logging.Log("OutputSystem", "Inbox Channel closed!")
	}

	return true
}

func (o *OutputSystem) ParseMessage(msg base.Message) bool {
	fmt.Println(msg.Content)
	return true
}
