package component

import (
	"log"

	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type curtain struct {
	img       *ebiten.Image
	rImg      *ebiten.Image
	x         float64
	y         float64
	direction direction
}

func NewCurtain(direction direction) Component {
	var img *ebiten.Image
	var err error
	switch direction {
	case helper.TOP:
		img, _, err = ebitenutil.NewImageFromFile("./assets/images/Stall/curtain_straight.png")
		if err != nil {
			log.Fatal(err)
		}
	case helper.RIGHT, helper.LEFT:
		img, _, err = ebitenutil.NewImageFromFile("./assets/images/Stall/curtain.png")
		if err != nil {
			log.Fatal(err)
		}
	default:
		panic("invalid direction")
	}

	rImg, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/curtain_rope.png")
	if err != nil {
		log.Fatal(err)
	}

	return &curtain{
		img:       img,
		rImg:      rImg,
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
	rOpt := &ebiten.DrawImageOptions{}
	rY := curtain.img.Bounds().Dy()/2 - curtain.rImg.Bounds().Dy()/2
	rX := curtain.rImg.Bounds().Dx() / 2

	switch direction {
	case helper.TOP:
		for x := 0.0; x < screenX; x += curtain.x {
			options = &ebiten.DrawImageOptions{}
			options.GeoM.Translate(x, 0)
			screen.DrawImage(curtain.img, options)
		}
		return nil
	case helper.RIGHT:
		options.GeoM.Scale(-1, 1)
		options.GeoM.Translate(screenX, defaultY)
		screen.DrawImage(curtain.img, options)

		rOpt.GeoM.Translate(screenX-float64(rX), defaultY+float64(rY))
		screen.DrawImage(curtain.rImg, rOpt)

	case helper.LEFT:
		options.GeoM.Translate(0, defaultY)
		screen.DrawImage(curtain.img, options)

		rOpt = options
		rOpt.GeoM.Translate(0-float64(rX), float64(rY))
		screen.DrawImage(curtain.rImg, rOpt)
	}

	return nil
}

func (curtain *curtain) Update() error {
	return nil
}

func (curtain *curtain) OnScreen() bool {
	return true
}
