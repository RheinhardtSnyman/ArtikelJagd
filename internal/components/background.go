package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type background struct {
	img *ebiten.Image
	x   float64
	y   float64
}

func NewBackground() Component {
	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/bg_green.png")
	if err != nil {
		log.Fatal(err)
	}
	return &background{
		img: img,
		x:   float64(img.Bounds().Dx()),
		y:   float64(img.Bounds().Dy()),
	}

}

func (background *background) Draw(screen *ebiten.Image) error {
	screenX := float64(screen.Bounds().Dx())
	screenY := float64(screen.Bounds().Dy())

	for x := 0.0; x < screenX; x += background.x {
		for y := 0.0; y < screenY; y += background.y {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(x, y)
			screen.DrawImage(background.img, options)
		}
	}

	return nil
}

func (d *background) Update() error {
	return nil
}

func (d *background) OnScreen() bool {
	return true
}
