package component

import "github.com/hajimehoshi/ebiten/v2"

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
