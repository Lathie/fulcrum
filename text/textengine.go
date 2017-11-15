package text

import (
//	"github.com/Lathie/fulcrum/base"
)

//TextEngine has the follow fields:
//Input - a InputSystem
//Logic - a LogicSystem
type TextEngine struct {
	//	Entities *EntitySystem
	Input *InputSystem
	Logic *LogicSystem
	//	Output   *OutputSystem
	Running bool
}

//NewEngine() creates a new TextEngine
func NewEngine() *TextEngine {
	MainLogic := NewLogicSystem()
	MainInput := NewInputSystem(MainLogic)
	Engine := TextEngine{Input: MainInput, Logic: MainLogic, Running: true}

	return &Engine
}

//Init creates a new engine
//This is deprecated, delete from interfaces/here
func (t *TextEngine) Init() {
	//	e.Entities = make(EntitySystem)
	//e.Input = make(InputSystem)
	//	e.Output = make(OutputSystem)
	//	e.Logic = make(LogicSystem)

	//	e.Input.Logic = &Logic
	//	e.Input.Init()

}

//Update() gets called each iteration of the gameloop
func (t *TextEngine) Update() bool {
	cont := true
	//t.Output.Update()
	t.Input.Update()
	//t.LogicSystem.Update()
	t.Logic.Update()
	//t.EntitySystem.Update()
	return cont
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
