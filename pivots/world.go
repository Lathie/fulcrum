package pivots

import (
//
)

const (
	MaxWorldSize   = 100
	MaxWorldHeight = 10
	RootX          = 50
	RootY          = 50
	RootZ          = 3
)

type World struct {
	Map [][][]Room
}

//Future: Store and retrieve worlds from a json encoding probably
func NewWorld() World {
	//Create inital root Room
	exits := [6]bool{true, false, false, false, false, false}
	name := "A empty room"
	description := "This is the initial room"
	descrip_all := "Upon further inspection, the room is completely empty. There is an exit to the north"
	location := [3]int{RootX, RootY, RootZ}
	newRoom := Room{Exits: exits, Name: name, Desc: description, DescAll: descrip_all, Location: location, Exists: true}

	otherExits := [6]bool{false, false, true, false, false, false}
	otherName := "Another empty room"
	otherDescription := "This is the second initial room"
	otherDescrip_all := "Upon further inspection, the room is also completely empty. There is an exit to the south"
	otherLocation := [3]int{RootX + 1, RootY, RootZ}
	otherNewRoom := Room{Exits: otherExits, Name: otherName, Desc: otherDescription, DescAll: otherDescrip_all, Location: otherLocation, Exists: true}

	rooms := make([][][]Room, MaxWorldSize)

	for i := range rooms {
		rooms[i] = make([][]Room, MaxWorldSize)
		for j := range rooms[i] {
			rooms[i][j] = make([]Room, MaxWorldHeight)
		}
	}
	world := World{Map: rooms}
	world.Map[RootX][RootY][RootZ] = newRoom
	world.Map[RootX+1][RootY][RootZ] = otherNewRoom
	return world
}
