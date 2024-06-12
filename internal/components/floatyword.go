package component

import (
	"bytes"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type image struct {
	img *ebiten.Image
	x   float64
	y   float64
}

type word struct {
	val      string
	fontSize float64
}

type floatyWord struct {
	img     image
	word    word
	aniX    animation
	aniY    animation
	x       float64
	y       float64
	show    bool
	score   *int
	variety int
	armed   *int
}

var faceSource *text.GoTextFaceSource

const (
	// Min background image to word size ratio
	// This makes sure word remains inside the background image
	minRatio = 1.9
	minY     = 85
	maxY     = 285
)

func init() {
	var err error
	faceSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
}

func getRandom(min, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}

func NewfloatyWord(score *int, aniX, aniY float64, variety int, armed *int) Component {

	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/cloud2.png")
	if err != nil {
		log.Fatal(err)
	}

	return &floatyWord{
		img: image{
			img: img,
			x:   float64(img.Bounds().Dx()),
			y:   float64(img.Bounds().Dy()),
		},
		word: word{
			val:      "Gemüẞe",
			fontSize: 22,
		},
		aniX: animation{
			tick:       0.0,
			speed:      float64(getRandom(80, 150)) / 100,
			changeSize: aniX,
			direction:  forward,
		},
		aniY: animation{
			tick:       0.0,
			speed:      float64(getRandom(10, 35)) / 100,
			changeSize: aniY,
			direction:  forward,
		},
		x:       -float64(img.Bounds().Dx()),
		y:       getRandom(minY, maxY),
		show:    true,
		score:   score,
		variety: variety,
		armed:   armed,
	}

}

func (floatyWord floatyWord) Draw(screen *ebiten.Image) error {
	// word
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   floatyWord.word.fontSize,
	}
	wordWidth, wordHeight := text.Measure(floatyWord.word.val, mplusNormalFont, 1)

	// background
	imgX := floatyWord.img.x
	imgY := floatyWord.img.y

	opImg := &ebiten.DrawImageOptions{}

	// Up scale background image to to min ratio to word width
	ratio := imgX / wordWidth
	if ratio < minRatio {
		scale := 1.0 + minRatio - ratio
		opImg.GeoM.Scale(scale, scale)
		imgX = imgX * scale
		imgY = imgY * scale
	}
	opImg.GeoM.Translate(floatyWord.x, floatyWord.y)

	// word
	// center of image
	wordX := floatyWord.x + imgX/2 - wordWidth/2
	wordY := floatyWord.y + imgY/2 - wordHeight/2

	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(wordX, wordY)
	opWord.ColorScale.ScaleWithColor(color.Black)

	// draw
	screen.DrawImage(floatyWord.img.img, opImg)
	text.Draw(screen, floatyWord.word.val, mplusNormalFont, opWord)

	return nil
}
func (floatyWord *floatyWord) Update() error {
	if floatyWord.x > floatyWord.aniX.changeSize+floatyWord.img.x {
		floatyWord.x = -floatyWord.img.x
		floatyWord.y = getRandom(minY, maxY)
	}
	floatyWord.x += floatyWord.aniX.speed

	if floatyWord.aniY.changeSize > 0 {
		floatyWord.y += floatyWord.aniY.speed * float64(floatyWord.aniY.direction)
		floatyWord.aniY.tick += floatyWord.aniY.speed * float64(floatyWord.aniY.direction)

		if floatyWord.aniY.tick >= floatyWord.aniY.changeSize && floatyWord.aniY.direction == forward {
			changeDirection(&floatyWord.aniY)
		} else if floatyWord.aniY.tick < 0 && floatyWord.aniY.direction == backwards {
			changeDirection(&floatyWord.aniY)
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		shot(floatyWord)
	}

	return nil
}

func shot(floatyWord *floatyWord) {
	boxMinX := floatyWord.x
	boxMaxX := boxMinX + floatyWord.img.x
	boxMinY := floatyWord.y
	boxMaxY := boxMinY + floatyWord.img.y

	x, y := ebiten.CursorPosition()
	if *floatyWord.armed != none {
		if x > int(boxMinX) && x < int(boxMaxX) && y > int(boxMinY) && y < int(boxMaxY) {
			// Got a hit
			if *floatyWord.armed == floatyWord.variety {
				*floatyWord.score++
				floatyWord.show = false
			} else {
				*floatyWord.score--
				floatyWord.x = -floatyWord.img.x
				floatyWord.y = getRandom(minY, maxY)
				if floatyWord.aniX.speed < 3 {
					floatyWord.aniX.speed += 2
				}
			}

		}
	}

}

func (floatyWord floatyWord) OnScreen() bool {
	return floatyWord.show
}
