package component

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Update(int) error
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
