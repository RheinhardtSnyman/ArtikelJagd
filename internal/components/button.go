package component

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type button struct {
	img      *ebiten.Image
	text     string
	fontSize int
	x        float64
	variety  int
	armed    *int
}

func NewButton(text string, x float64, variety int, armed *int) Component {

	img, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/button.png")
	if err != nil {
		log.Fatal(err)
	}
	return &button{
		img:      img,
		text:     text,
		fontSize: 30,
		x:        x,
		variety:  variety,
		armed:    armed,
	}
}

func getColor(variety int) color.Color {
	switch variety {
	case red:
		return color.RGBA{R: 195, G: 0, B: 50, A: 255}
	case blue:
		return color.RGBA{R: 0, G: 5, B: 200, A: 255}
	case yellow:
		return color.RGBA{R: 195, G: 170, B: 0, A: 255}
	}
	return color.RGBA{R: 255, G: 255, B: 255, A: 255}
}

func (button button) Draw(screen *ebiten.Image) error {
	//set text options
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   float64(button.fontSize),
	}
	wordWidth, wordHeight := text.Measure(button.text, mplusNormalFont, 1)
	btnY := float64(screen.Bounds().Dy()) - 60 - wordHeight/2

	//draw btn
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(button.x-float64(button.img.Bounds().Dx()/2)+wordWidth/2, btnY-float64(button.img.Bounds().Dy()/2)+wordHeight/2)
	screen.DrawImage(button.img, options)

	//draw text
	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(button.x, btnY)
	opWord.ColorScale.ScaleWithColor(getColor(button.variety))

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
