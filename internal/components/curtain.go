package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type curtain struct {
	img       *ebiten.Image
	x         float64
	y         float64
	direction direction
}

const (
	north direction = iota
	east
	west
)

func NewCurtain(direction direction) Component {
	var img *ebiten.Image
	var err error
	switch direction {
	case north:
		img, _, err = ebitenutil.NewImageFromFile("./assets/images/Stall/curtain_straight.png")
		if err != nil {
			log.Fatal(err)
		}
	case east, west:
		img, _, err = ebitenutil.NewImageFromFile("./assets/images/Stall/curtain.png")
		if err != nil {
			log.Fatal(err)
		}
	default:
		panic("invalid direction")
	}

	return &curtain{
		img:       img,
		x:         float64(img.Bounds().Dx()),
		y:         float64(img.Bounds().Dy()),
		direction: direction,
	}
}

func (curtain *curtain) Draw(screen *ebiten.Image) error {
	defaultY := 55.0
	direction := curtain.direction
	screenX := float64(screen.Bounds().Dx())
	options := &ebiten.DrawImageOptions{}

	switch direction {
	case north:
		for x := 0.0; x < screenX; x += curtain.x {
			options = &ebiten.DrawImageOptions{}
			options.GeoM.Translate(x, 0)
			screen.DrawImage(curtain.img, options)
		}
		return nil
	case east:
		options.GeoM.Scale(-1, 1)
		options.GeoM.Translate(screenX, defaultY)
	case west:
		options.GeoM.Translate(0, defaultY)
	}

	screen.DrawImage(curtain.img, options)

	return nil
}

func (curtain *curtain) Update() error {
	return nil
}

func (curtain *curtain) OnScreen() bool {
	return true
}
