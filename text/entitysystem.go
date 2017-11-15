package text

import "github.com/Lathie/fulcrum/base"

/*
Entity system handles updating all the entities in the game
*/
type EntitySystem struct {
	//Entities []base.Entity
}

func (e *EntitySystem) Update() {
	//	for _, v := range e.Entities {
	//		v.Update()
	//	}
}

func (e *EntitySystem) AddEntity(ent *base.Entity) {
	//	e.Entities = append(e.Entities, e)
}
