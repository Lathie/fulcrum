package base

//Entities must contain the EntityComponents struct
type EntityComponents struct {
	Components []Component
}

//Entities must implement the Update() function
type Entity interface {
	Update()
}
