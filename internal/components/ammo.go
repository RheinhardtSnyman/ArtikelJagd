package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	empty = iota
	half
	full
	super
)

const bulletCount = 3

type ammo struct {
	emptyImg *ebiten.Image
	fullImg  *ebiten.Image
	halfImg  *ebiten.Image
	superImg *ebiten.Image
	x        float64
	y        float64
	bullets  [bulletCount]int
	ammo     int
}

func NewAmmo(x float64, y float64) Component {
	emptyImg, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/icon_bullet_empty_long.png")
	if err != nil {
		log.Fatal(err)
	}

	fullImg, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/icon_bullet_silver_long.png")
	if err != nil {
		log.Fatal(err)
	}

	halfImg, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/icon_bullet_silver_short.png")
	if err != nil {
		log.Fatal(err)
	}

	superImg, _, err := ebitenutil.NewImageFromFile("./assets/images/HUD/icon_bullet_gold_long.png")
	if err != nil {
		log.Fatal(err)
	}

	var blts [bulletCount]int

	for i := 0; i < bulletCount; i++ {
		blts[i] = full
	}

	return &ammo{
		emptyImg: emptyImg,
		fullImg:  fullImg,
		halfImg:  halfImg,
		superImg: superImg,
		x:        x,
		y:        y,
		bullets:  blts,
		ammo:     bulletCount * 2,
	}
}

func refreshAmmo(ammo *ammo) {
	blt := [bulletCount]int{0, 0, 0}
	for idx := 0; idx < ammo.ammo; idx++ {
		for i := 0; i < bulletCount && idx < ammo.ammo; i++ {
			blt[i] += 1
			idx++
		}
	}
	ammo.bullets = blt
}

func (ammo *ammo) Draw(screen *ebiten.Image) error {

	img := ammo.emptyImg

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(ammo.x, ammo.y)

	for i := 0; i < bulletCount; i++ {
		switch ammo.bullets[i] {
		case empty:
		case half:
			img = ammo.halfImg
		case full:
			img = ammo.fullImg
		case super:
			img = ammo.superImg
		}
		options.GeoM.Translate(float64(img.Bounds().Dx())+7.5, 0)
		screen.DrawImage(img, options)
	}
	return nil
}

func (ammo *ammo) Update() error {
	return nil
}

func (ammo *ammo) OnScreen() bool {
	return true
}
