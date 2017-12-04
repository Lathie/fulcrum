package systems

import (
	"fmt"
	"github.com/Lathie/fulcrum/base"
	"github.com/Lathie/fulcrum/logging"
	"sync"
)

//Constants
//MessageLimit - how many messages are allowed in a system
const (
	MessageLimit = 10
)

//TextEngine has the follow fields:
//Input - a InputSystem
//Logic - a LogicSystem
type TextEngine struct {
	Input *InputSystem
	Logic *LogicSystem
	World *WorldSystem
	Bus   *MessageBus

	Running bool
	Debug   bool
}

//NewEngine() creates a new TextEngine
func NewEngine(debug bool) *TextEngine {
	BusInbox := make(chan base.Message, MessageLimit)
	InputInbox := make(chan base.Message, MessageLimit)
	LogicInbox := make(chan base.Message, MessageLimit)
	WorldInbox := make(chan base.Message, MessageLimit)

	MainLogic := NewLogicSystem(LogicInbox, BusInbox)
	MainInput := NewInputSystem(InputInbox, BusInbox)
	MainWorld := NewWorldSystem(WorldInbox, BusInbox)
	MainBus := NewMessageBus(BusInbox, InputInbox, LogicInbox, WorldInbox, debug)

	Engine := TextEngine{Input: MainInput, Logic: MainLogic, World: MainWorld, Bus: MainBus, Running: true, Debug: debug}

	logging.Log("Master", "Engine Initialized")

	return &Engine
}

//Init creates a new engine
//This is deprecated, delete from interfaces/here
func (t *TextEngine) Init() {
	//	e.Entities = make(EntitySystem)
	//  e.Input = make(InputSystem)
	//	e.Output = make(OutputSystem)
	//	e.Logic = make(LogicSystem)

	//	e.Input.Logic = &Logic
	//	e.Input.Init()

}

//Update() gets called each iteration of the gameloop
//deprecated for now
func (t *TextEngine) Update() bool {

	//t.Input.Update()
	//t.Bus.Update()
	//t.Logic.Update()

	return true
}

//MainLoop() is the main loop of the game engine.
//It spawns goroutines and then waits for them to finish
func (t *TextEngine) MainLoop() {
	fmt.Printf("Main loop starting \n")

	var wg sync.WaitGroup
	wg.Add(3)

	go t.Launch(t.Input, "Input")
	go t.Launch(t.Logic, "Logic")
	go t.Launch(t.Bus, "Message Bus")
	go t.Launch(t.World, "World")

	wg.Wait()
}

//Launch takes an interface and wraps it's Update() function in a for loop
//This is to start a goroutine for a system
func (t *TextEngine) Launch(s base.System, str string) {
	text := str + " thread started"
	logging.Log("Master", text)
	for {
		s.Update()
	}
}
