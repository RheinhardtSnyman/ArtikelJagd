package component

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type button struct {
	img      *ebiten.Image
	imgPress *ebiten.Image
	imgHov   *ebiten.Image
	text     string
	fontSize int
	x        float64
	variety  int
	armed    *int
	scale    float64
}

func NewButton(text string, x float64, variety int, armed *int) Component {

	img, imgPress, imgHov := getImages(variety)
	return &button{
		img:      img,
		imgPress: imgPress,
		imgHov:   imgHov,
		text:     text,
		fontSize: 30,
		x:        x,
		variety:  variety,
		armed:    armed,
		scale:    0.75,
	}
}

func getImages(variety int) (*ebiten.Image, *ebiten.Image, *ebiten.Image) {
	color := ""
	switch variety {
	case red:
		color = "_red"
	case blue:
		color = "_blue"
	case green:
		color = "_green"
	}

	imgs := [3]*ebiten.Image{}

	imgUrls := [3]string{fmt.Sprintf("./assets/images/HUD/button%s.png", color),
		fmt.Sprintf("./assets/images/HUD/button%s_press.png", color),
		fmt.Sprintf("./assets/images/HUD/button%s_hov.png", color)}

	for idx, strUrl := range imgUrls {
		var err error
		imgs[idx], _, err = ebitenutil.NewImageFromFile(strUrl)
		if err != nil {
			log.Fatal(err)
		}
	}

	return imgs[0], imgs[1], imgs[2]
}

func (button button) Draw(screen *ebiten.Image) error {
	//set text options
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   float64(button.fontSize),
	}
	wordWidth, wordHeight := text.Measure(button.text, mplusNormalFont, 1)
	btnY := float64(screen.Bounds().Dy()) - 60

	//Button and mouse positions
	btnImage := button.img
	x, y := ebiten.CursorPosition()
	boxMinX := button.x - float64(btnImage.Bounds().Dx()/2) + wordWidth/2
	boxMaxX := boxMinX + float64(btnImage.Bounds().Dx())
	boxMinY := btnY - float64(btnImage.Bounds().Dy()/2)
	boxMaxY := btnY + float64(btnImage.Bounds().Dy()/2)

	//If variety selected show only pressed btn
	if *button.armed == button.variety {
		btnImage = button.imgPress
	} else {
		if x > int(boxMinX) && x < int(boxMaxX) && y > int(boxMinY) && y < int(boxMaxY) {
			btnImage = button.imgHov
		}
	}

	//select variety on btn click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if x > int(boxMinX) && x < int(boxMaxX) && y > int(boxMinY) && y < int(boxMaxY) {
			*button.armed = button.variety
		}
	}

	//draw btn
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(boxMinX, boxMinY)
	screen.DrawImage(btnImage, options)

	//draw text
	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(button.x, btnY-wordHeight/2)
	opWord.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})

	text.Draw(screen, button.text, mplusNormalFont, opWord)

	return nil
}

func (button *button) Update() error {
	return nil
}

func (button *button) OnScreen() bool {
	return true
}
