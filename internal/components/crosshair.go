package component

import (
	"log"

	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type crosshair struct {
	img   *ebiten.Image
	x     float64
	y     float64
	armed *int
}

func getImg(armed int) *ebiten.Image {

	name := "crosshair_white_small"
	switch armed {
	case helper.RED:
		name = "crosshair_red_small"
	case helper.BLUE:
		name = "crosshair_blue_small"
	case helper.GREEN:
		name = "crosshair_green_small"
	}
	if armed != helper.NONE {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}
	img, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/" + name + ".png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func NewCrosshair(armed *int) Component {
	img := getImg(*armed)
	return &crosshair{
		img:   img,
		x:     float64(img.Bounds().Dx()),
		y:     float64(img.Bounds().Dy()),
		armed: armed,
	}
}

func (crosshair *crosshair) Draw(screen *ebiten.Image) error {

	if *crosshair.armed != helper.NONE {
		crosshair.img = getImg(*crosshair.armed)
		curX, curY := ebiten.CursorPosition()

		x := float64(curX) - crosshair.x/2
		y := float64(curY) - crosshair.y/2

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(x, y)
		screen.DrawImage(crosshair.img, options)
	}

	return nil
}

func (crosshair *crosshair) Update() error {
	return nil
}

func (crosshair *crosshair) OnScreen() bool {
	return true
}
