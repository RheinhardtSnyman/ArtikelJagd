package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var img *ebiten.Image

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	deskH := 100

	border := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	border.Fill(color.RGBA{0x80, 0x57, 0x2e, 0xff})

	// screen.DrawImage(img, nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(screen.Bounds().Dy()-deskH))
	opb := &ebiten.DrawImageOptions{}
	opb.GeoM.Translate(0, float64(screen.Bounds().Dy()-4-deskH))

	screen.DrawImage(border, opb)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(800, 480)
	ebiten.SetWindowTitle("Hello, World!")

	var err error
	img, _, err = ebitenutil.NewImageFromFile("./assets/images/Stall/bg_wood.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
