package text

import "github.com/Lathie/fulcrum/base"

type TextEngine struct {
	Entities *EntitySystem
	Input    *InputSystem
}

func (e *TextEngine) Init() {

}

func (e *TextEngine) Update() {
	e.InputSystem.Update()
	//e.LogicSystem.Update()
	e.EntitySystem.Update()
}

/*
The game loop should:
Print Prompt
Get Line
Process Game Logic
Update all states
(repeat)
*/
func (e *TextEngine) MainLoop() {
	for {
		e.Update()
	}
}
