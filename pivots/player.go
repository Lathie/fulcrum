package pivots

import (
//
)

type Player struct {
	Location [3]int
	Sudo     bool
}

func NewPlayer(sudo bool) Player {
	loc := [3]int{RootX, RootY, RootZ}
	player := Player{Location: loc, Sudo: sudo}

	return player

}
