package pivots

import (
//	"github.com/Lathie/fulcrum/systems"
)

type Room struct {
	Exits    [6]bool
	Name     string
	Desc     string
	DescAll  string
	Location [3]int
	Exists   bool
}
