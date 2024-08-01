package component

import (
	"image/color"

	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type menu struct {
	text     string
	fontSize int
}

func NewMenu(text string) Component {
	return &menu{
		text:     text,
		fontSize: helper.FONT_LARGE,
	}
}

func (menu menu) Draw(screen *ebiten.Image) error {
	// Draw text
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   float64(menu.fontSize),
	}
	cntX := float64(screen.Bounds().Dx()) / 2
	cntY := float64(screen.Bounds().Dy()) / 2
	wordWidth, wordHeight := text.Measure(menu.text, mplusNormalFont, 1)

	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(cntX-wordWidth/2, cntY-wordHeight/2)
	opWord.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, menu.text, mplusNormalFont, opWord)
	return nil
}

func (menu *menu) Update() error {
	return nil
}

func (menu *menu) OnScreen() bool {
	return true
}
