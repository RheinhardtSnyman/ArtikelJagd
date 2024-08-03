package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type empty struct {
}

func NewEmpty() Component {

	return &empty{}
}

func (empty *empty) Draw(screen *ebiten.Image) error {
	return nil
}

func (empty *empty) Update() error {
	return nil
}

func (empty *empty) OnScreen() bool {
	return true
}
