package pivots

import (
//	"github.com/Lathie/fulcrum/systems"
)

const (
	North = iota
	East
	South
	West
	Up
	Down
)

const (
	X = iota
	Y
	Z
)

type Room struct {
	Exits    [6]int
	Name     string
	Desc     string
	DescAll  string
	Location [3]int
	Exists   bool
}
