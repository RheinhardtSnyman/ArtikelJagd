package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type mountian struct {
	imgs  [2]*ebiten.Image
	y     float64
	combo []int
}

func getRange(w int, count int, imgW int) []int {
	imgCombo := []int{}
	for x := 0; x < w; x += imgW {
		imgCombo = append(imgCombo, int(getRandom(0, count)))
	}
	return imgCombo
}

func NewMountian(y float64, w int) Component {
	imgUrls := []string{"./assets/images/Stall/grass1.png",
		"./assets/images/Stall/grass2.png"}

	imgs := [2]*ebiten.Image{}
	for idx, strUrl := range imgUrls {
		var err error
		imgs[idx], _, err = ebitenutil.NewImageFromFile(strUrl)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &mountian{
		imgs:  imgs,
		y:     y,
		combo: getRange(w, len(imgUrls), 132),
	}
}

func (mountian *mountian) Draw(screen *ebiten.Image) error {
	screenX := screen.Bounds().Dx()

	idx := 0
	curImg := mountian.imgs[mountian.combo[idx]]

	for x := 0; x < screenX; x += curImg.Bounds().Dx() {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x), mountian.y)
		screen.DrawImage(curImg, options)
		curImg = mountian.imgs[mountian.combo[idx]]
		idx++
	}

	return nil
}

func (mountian *mountian) Update() error {
	return nil
}

func (mountian *mountian) OnScreen() bool {
	return true
}
