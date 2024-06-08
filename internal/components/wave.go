package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type animate struct {
	tick       float64
	speed      float64
	changeSize float64
	dirForward bool
}

type wave struct {
	img   *ebiten.Image
	x     float64
	y     float64
	posX  float64
	posY  float64
	aniDr bool
	aniX  animate
	aniY  animate
}

func NewWave(onewayAniX bool, waveName string, optXY ...float64) Component {
	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/" + waveName + ".png")
	if err != nil {
		log.Fatal(err)
	}

	// defaults
	posX := 0.0
	aniX := animate{
		tick:       0,
		speed:      0.2,
		changeSize: 32,
		dirForward: true,
	}
	posY := 200.0
	aniY := animate{
		tick:       0,
		speed:      0.1,
		changeSize: 8,
		dirForward: true,
	}

	// optXY override default
	if len(optXY) > 0 {
		aniXY := []*float64{
			&posX,
			&aniX.speed,
			&aniX.changeSize,
			&posY,
			&aniY.speed,
			&aniY.changeSize,
		}

		for i, val := range optXY {
			if val > 0 {
				*aniXY[i] = val
			}
		}
	}
	return &wave{
		img:   img,
		x:     float64(img.Bounds().Dx()),
		y:     float64(img.Bounds().Dy()),
		posX:  posX,
		posY:  posY,
		aniDr: onewayAniX,
		aniX:  aniX,
		aniY:  aniY,
	}
}

func (wave *wave) Draw(screen *ebiten.Image) error {
	screenX := float64(screen.Bounds().Dx())
	screenY := float64(screen.Bounds().Dy())

	for x := -wave.x; x < screenX; x += wave.x {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(x+wave.posX+wave.aniX.tick, screenY-wave.posY+wave.aniY.tick)
		screen.DrawImage(wave.img, options)
	}

	return nil
}

func (wave *wave) Update(tick int) error {

	if wave.aniDr {
		if wave.aniX.tick >= wave.x {
			wave.aniX.tick = 0
		}
		wave.aniX.tick += wave.aniX.speed
	} else {
		aniUpDownCntr(&wave.aniX)
	}
	aniUpDownCntr(&wave.aniY)

	return nil
}

func aniUpDownCntr(ani *animate) {
	if ani.dirForward {
		ani.tick += ani.speed
		if ani.tick >= ani.changeSize {
			ani.dirForward = false
		}
	} else {
		ani.tick -= ani.speed
		if ani.tick <= 0 {
			ani.dirForward = true
		}
	}
}

func (wave *wave) OnScreen() bool {
	return true
}
