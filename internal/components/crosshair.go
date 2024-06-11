package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type crosshair struct {
	img *ebiten.Image
	x   float64
	y   float64
}

func NewCrosshair() Component {
	img, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/crosshair_white_small.png")
	if err != nil {
		log.Fatal(err)
	}
	return &crosshair{
		img: img,
		x:   float64(img.Bounds().Dx()),
		y:   float64(img.Bounds().Dy()),
	}
}

func (crosshair *crosshair) Draw(screen *ebiten.Image) error {

	curX, curY := ebiten.CursorPosition()

	x := float64(curX) - crosshair.x/2
	y := float64(curY) - crosshair.y/2

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(x, y)
	screen.DrawImage(crosshair.img, options)

	return nil
}

func (crosshair *crosshair) Update() error {
	return nil
}

func (crosshair *crosshair) OnScreen() bool {
	return true
}
