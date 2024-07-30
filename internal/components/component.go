// This is the general component interface file
// TODO move helper functions to a different file
package component

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	Update() error
	Draw(*ebiten.Image) error
	OnScreen() bool
}

type direction int

const (
	forward   direction = 1
	backwards direction = -1
)

// iota declared this way will result in values of 0,1,2,3 and so on in underlying variables
const (
	red = iota
	blue
	green
	none
)

type animation struct {
	tick       float64
	speed      float64
	changeSize float64
	direction  direction
}

func changeDirection(ani *animation) {
	ani.direction = ani.direction * -1
}

// TODO this func can posibly reused
func getRandom(min, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}
