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

type animation struct {
	tick       float64
	speed      float64
	changeSize float64
	dirForward bool
}

type floatyWord struct {
	img  image
	word word
	aniX animation
	aniY animation
	x    float64
	y    float64
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

func NewfloatyWord(aniX, aniY float64) Component {

	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/cloud1.png")
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
			speed:      1.0,
			changeSize: aniX,
			dirForward: true,
		},
		aniY: animation{
			tick:       0.0,
			speed:      0.1,
			changeSize: aniY,
			dirForward: true,
		},
		x: -float64(img.Bounds().Dx()),
		y: getRandom(minY, maxY),
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
func (floatyWord *floatyWord) Update(tick int) error {
	if floatyWord.x > floatyWord.aniX.changeSize+floatyWord.img.x {
		floatyWord.x = -floatyWord.img.x
		floatyWord.y = getRandom(minY, maxY)
	}
	floatyWord.x += floatyWord.aniX.speed

	if floatyWord.aniY.changeSize > 0 {
		if floatyWord.aniY.tick < floatyWord.aniY.changeSize && floatyWord.aniY.dirForward {
			floatyWord.y += floatyWord.aniY.speed
			floatyWord.aniY.tick += floatyWord.aniY.speed
		} else {
			floatyWord.aniY.dirForward = false
		}

		if floatyWord.aniY.tick > 0 && !floatyWord.aniY.dirForward {
			floatyWord.y -= floatyWord.aniY.speed
			floatyWord.aniY.tick -= floatyWord.aniY.speed
		} else {
			floatyWord.aniY.dirForward = true
		}

	}

	return nil
}
func (floatyWord floatyWord) OnScreen() bool {
	return true
}
