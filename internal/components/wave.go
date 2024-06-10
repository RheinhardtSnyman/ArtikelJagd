package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type wave struct {
	img   *ebiten.Image
	x     float64
	y     float64
	posX  float64
	posY  float64
	aniDr bool
	aniX  animation
	aniY  animation
}

func NewWave(onewayAniX bool, waveName string, optXY ...float64) Component {
	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/" + waveName + ".png")
	if err != nil {
		log.Fatal(err)
	}

	// defaults
	posX := 0.0
	aniX := animation{
		tick:       0,
		speed:      0.2,
		changeSize: 32,
		direction:  forward,
	}
	posY := 200.0
	aniY := animation{
		tick:       0,
		speed:      0.1,
		changeSize: 8,
		direction:  forward,
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

func aniUpDownCntr(ani *animation) {
	ani.tick += ani.speed * float64(ani.direction)
	if ani.direction == forward && ani.tick >= ani.changeSize {
		changeDirection(ani)
	} else if ani.direction == backwards && ani.tick <= 0 {
		changeDirection(ani)
	}
}

func (wave *wave) OnScreen() bool {
	return true
}
