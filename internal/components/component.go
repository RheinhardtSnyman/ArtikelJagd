// This is the general component interface file
// TODO move helper functions to a different file
package component

import (
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

type animation struct {
	tick       float64
	speed      float64
	changeSize float64
	direction  direction
}

func changeDirection(ani *animation) {
	ani.direction = ani.direction * -1
}
