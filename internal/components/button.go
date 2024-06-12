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
	variety  int
	armed    *int
}

func NewButton(text string, x float64, variety int, armed *int) Component {
	return &button{
		text:     text,
		fontSize: 30,
		x:        x,
		variety:  variety,
		armed:    armed,
	}
}

func (button button) Draw(screen *ebiten.Image) error {
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   float64(button.fontSize),
	}
	wordWidth, wordHeight := text.Measure(button.text, mplusNormalFont, 1)
	btnY := float64(screen.Bounds().Dy()) - 60 - wordHeight/2

	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(button.x, btnY)
	opWord.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, button.text, mplusNormalFont, opWord)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		boxMinX := button.x
		boxMaxX := button.x + wordWidth
		boxMinY := btnY
		boxMaxY := boxMinY + wordHeight

		x, y := ebiten.CursorPosition()

		if x > int(boxMinX) && x < int(boxMaxX) && y > int(boxMinY) && y < int(boxMaxY) {
			*button.armed = button.variety
		}
	}

	return nil
}

func (button *button) Update() error {
	return nil
}

func (button *button) OnScreen() bool {
	return true
}
