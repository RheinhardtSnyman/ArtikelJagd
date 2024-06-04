package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type table struct {
	img *ebiten.Image
	x   float64
	y   float64
}

func NewTable() Component {
	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/bg_wood.png")
	if err != nil {
		log.Fatal(err)
	}

	return &table{
		img: img,
		x:   float64(img.Bounds().Dx()),
		y:   float64(img.Bounds().Dy()),
	}

}

func (table *table) Draw(trgt *ebiten.Image) error {

	options := &ebiten.DrawImageOptions{}
	trgt.DrawImage(table.img, options)

	return nil
}

func (d *table) Update(_ *ebiten.Image, _ uint) error {
	return nil
}

func (d *table) OnScreen() bool {
	return true
}
