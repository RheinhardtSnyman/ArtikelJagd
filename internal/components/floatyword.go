package component

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type floatyWord struct {
	img  *ebiten.Image
	word string
	size float64
	x    float64
	y    float64
}

var FaceSource *text.GoTextFaceSource

func init() {
	var err error
	FaceSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
}
func NewfloatyWord() Component {

	img, _, err := ebitenutil.NewImageFromFile("./assets/images/Stall/cloud1.png")
	if err != nil {
		log.Fatal(err)
	}

	return &floatyWord{
		img:  img,
		word: "Gemüẞe",
		size: 22,
		x:    float64(img.Bounds().Dx()),
		y:    float64(img.Bounds().Dy()),
	}

}

func (floatyWord floatyWord) Draw(screen *ebiten.Image) error {

	x := 0.0
	y := 285.0 //85 to 285

	// word
	mplusNormalFont := &text.GoTextFace{
		Source: FaceSource,
		Size:   floatyWord.size,
	}
	wordWidth, wordHeight := text.Measure(floatyWord.word, mplusNormalFont, 1)

	// background
	imgX := floatyWord.x
	imgY := floatyWord.y

	opImg := &ebiten.DrawImageOptions{}

	// Up scale background image to to min ratio to word width
	minRatio := 2.0
	ratio := imgX / wordWidth
	if ratio < minRatio {
		scale := 1.0 + minRatio - ratio
		opImg.GeoM.Scale(scale, scale)
		imgX = imgX * scale
		imgY = imgY * scale
	}
	opImg.GeoM.Translate(x, y)

	// word
	// center of image
	wordX := x + imgX/2 - wordWidth/2
	wordY := y + imgY/2 - wordHeight/2

	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(wordX, wordY)
	opWord.ColorScale.ScaleWithColor(color.Black)

	// draw
	screen.DrawImage(floatyWord.img, opImg)
	text.Draw(screen, floatyWord.word, mplusNormalFont, opWord)

	return nil
}
func (floatyWord floatyWord) Update(_ int) error {
	return nil
}
func (floatyWord floatyWord) OnScreen() bool {
	return true
}
