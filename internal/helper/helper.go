package helper

import (
	"math/rand"
)

// iota declared this way will result in values of 0,1,2,3 and so on in underlying variables
const (
	NONE = iota
	RED
	BLUE
	GREEN
	TOP
	LEFT
	RIGHT
	BOTTOM
)

var DemoMode = false

var FONT_LARGE = 110
var TITLE_TEXT = "ArtikelJagd"

var START_TEXT = "Los geht's!"
var END_TEXT = "Aua, vorbei!"

var BTN_TEXT = map[int]string{
	RED:   "Die",
	BLUE:  "Der",
	GREEN: "Das",
}

func GetRandom(min, max int) float64 { //*Note: shared func are required to start with Capital letters
	return float64(rand.Intn(max-min) + min)
}
