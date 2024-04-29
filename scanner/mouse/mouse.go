package mouse

import (
	"github.com/go-vgo/robotgo"
)

func Location() (int, int) {
	x, y := robotgo.Location()
	return x, y
}
