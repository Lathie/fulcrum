package text

import "github.com/Lathie/fulcrum/base"

//Entity system handles updating all the entities in the game
type EntitySystem struct {
	//Entities []base.Entity
}

//Update() is called each game iteration
func (e *EntitySystem) Update() {
	//	for _, v := range e.Entities {
	//		v.Update()
	//	}
}

//AddEntity allows entities to be added to the game
func (e *EntitySystem) AddEntity(ent *base.Entity) {
	//	e.Entities = append(e.Entities, e)
}
