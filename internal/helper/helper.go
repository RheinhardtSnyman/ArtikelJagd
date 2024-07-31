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

func GetRandom(min, max int) float64 { //*Note: shared func are required to start with Capital letters
	return float64(rand.Intn(max-min) + min)
}
