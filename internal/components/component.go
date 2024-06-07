package component

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Update(int) error
	Draw(*ebiten.Image) error
	OnScreen() bool
}
