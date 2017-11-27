package systems

import (
	"github.com/Lathie/fulcrum/base"
)

const (
	MessageLimit = 50
)

//TextEngine has the follow fields:
//Input - a InputSystem
//Logic - a LogicSystem
type TextEngine struct {
	Input   *InputSystem
	Logic   *LogicSystem
	Bus     *MessageBus
	Running bool
	Debug   bool
}

//NewEngine() creates a new TextEngine
func NewEngine(debug bool) *TextEngine {
	BusInbox := make(chan base.Message, MessageLimit)
	InputInbox := make(chan base.Message, MessageLimit)
	LogicInbox := make(chan base.Message, MessageLimit)

	MainLogic := NewLogicSystem(LogicInbox, BusInbox)
	MainInput := NewInputSystem(InputInbox, BusInbox)
	MainBus := NewMessageBus(BusInbox, InputInbox, LogicInbox, debug)
	Engine := TextEngine{Input: MainInput, Logic: MainLogic, Bus: MainBus, Running: true, Debug: debug}

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
func (t *TextEngine) Update() bool {

	t.Input.Update()
	t.Bus.Update()
	t.Logic.Update()

	return true
}

//MainLoop() is the main loop of the game engine. It works as follows:
//The game loop should:
//Print Prompt
//Get Line
//Process Game Logic
//Update all states
//(repeat)
//
func (t *TextEngine) MainLoop() {
	for t.Running {
		t.Running = t.Update()
	}
}
