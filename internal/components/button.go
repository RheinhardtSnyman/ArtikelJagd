package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type button struct {
	text     string
	fontSize int
	x        float64
}

func NewButton(text string, x float64) Component {
	return &button{
		text:     text,
		fontSize: 30,
		x:        x,
	}
}

func (button button) Draw(screen *ebiten.Image) error {
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   float64(button.fontSize),
	}
	_, wordHeight := text.Measure(button.text, mplusNormalFont, 1)
	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(button.x, float64(screen.Bounds().Dy())-60-wordHeight/2)
	opWord.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, button.text, mplusNormalFont, opWord)

	return nil
}

func (button *button) Update() error {
	return nil
}

func (button *button) OnScreen() bool {
	return true
}
