package main

import (
	"github.com/Lathie/fulcrum/systems"
)

func main() {
	engine := systems.NewEngine(true) //Create a new engine in debug mode
	engine.MainLoop()
}
