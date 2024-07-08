package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ammo struct {
	eagleImg *ebiten.Image
	padding  int
	x        float64
	y        float64
	total    *int
}

func NewAmmo(x float64, y float64, padding int, total *int) Component {
	eagleImg, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/eagle_small.png")
	if err != nil {
		log.Fatal(err)
	}

	return &ammo{
		eagleImg: eagleImg,
		padding:  padding,
		x:        x,
		y:        y,
		total:    total,
	}
}

func (ammo *ammo) Draw(screen *ebiten.Image) error {

	img := ammo.eagleImg

	posX := screen.Bounds().Dx() - (img.Bounds().Dx()+ammo.padding)**ammo.total

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(posX), ammo.y)

	for i := 0; i < *ammo.total; i++ {
		screen.DrawImage(img, options)
		options.GeoM.Translate(float64(img.Bounds().Dx()+ammo.padding), 0)
	}
	return nil
}

func (ammo *ammo) Update() error {
	return nil
}

func (ammo *ammo) OnScreen() bool {
	return true
}
