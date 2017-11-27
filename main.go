package main

import (
	"github.com/Lathie/fulcrum/text"
)

func main() {
	engine := text.NewEngine(true) //Create a new engine in debug mode
	engine.MainLoop()
}
