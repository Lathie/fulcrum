package base

type EntityComponents struct{
	Components []Component
}

type Entity interface {
	Update()
}
