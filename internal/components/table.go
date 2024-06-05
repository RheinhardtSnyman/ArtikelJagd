package component

import (
	"image/color"
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

func (table *table) Draw(screen *ebiten.Image) error {
	defaultY := 120
	screenX := float64(screen.Bounds().Dx())
	tableY := float64(screen.Bounds().Dy() - defaultY)
	defaultBorder := 4

	// Table border - first
	border := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	border.Fill(color.RGBA{0x80, 0x57, 0x2e, 0xff})
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, tableY-float64(defaultBorder))
	screen.DrawImage(border, options)

	// Table - ontop of border
	for x := 0.0; x < screenX; x += table.x {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(x, tableY)
		screen.DrawImage(table.img, options)
	}

	return nil
}

func (d *table) Update(_ *ebiten.Image, _ uint) error {
	return nil
}

func (d *table) OnScreen() bool {
	return true
}
