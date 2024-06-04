package component

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Update(*ebiten.Image, uint) error
	Draw(*ebiten.Image) error
	OnScreen() bool
}

type direction int
