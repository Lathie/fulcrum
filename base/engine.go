package base

/*
What does the Engine need: a TODO list
Struct:
It needs a list of systems

Interface:
It needs a update
It needs a MainLoop
It needs a way of adding systems

Interactions between systems and components:

So components will be interfaces to systems

----You should create a component by passing a pointer to a system and then
--------leverage the utilities of the system in order to use the component
----So how do entities get updated?

----There will be an entity manager system probably that will loop through the system and
--------call update on all the entities which then will call update on the components

TODO:
Better Message Bus
*/

//Engines must implement Init(), Update(), and MainLoop()
type Engine interface {
	Init()
	Update() bool
	MainLoop()
}
