package text

import (
	"github.com/Lathie/fulcrum/base"
)

type TextEngine struct {
	//	Entities *EntitySystem
	Input *InputSystem
	Logic *LogicSystem
	//	Output   *OutputSystem
	Running bool
}

func NewEngine() *TextEngine {
	MainLogic := LogicSystem{Messages: make([]base.Message, 0)}
	MainInput := NewInputSystem(&MainLogic)
	Engine := TextEngine{Input: MainInput, Logic: &MainLogic, Running: true}

	return &Engine

}

func (t *TextEngine) Init() {
	//	e.Entities = make(EntitySystem)
	//e.Input = make(InputSystem)
	//	e.Output = make(OutputSystem)
	//	e.Logic = make(LogicSystem)

	//	e.Input.Logic = &Logic
	//	e.Input.Init()

}

func (t *TextEngine) Update() bool {
	cont := true
	//t.Output.Update()
	t.Input.Update()
	//t.LogicSystem.Update()
	t.Logic.Update()
	//t.EntitySystem.Update()
	return cont
}

/*
The game loop should:
Print Prompt
Get Line
Process Game Logic
Update all states
(repeat)
*/
func (t *TextEngine) MainLoop() {
	for t.Running {
		t.Running = t.Update()
	}
}
