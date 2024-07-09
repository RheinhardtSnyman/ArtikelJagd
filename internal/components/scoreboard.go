package component

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type scoreboard struct {
	text         string
	fontSize     int
	score        *int
	lives        *int
	bonusEagleAt int
}

func NewScoreboard(lives *int, score *int) Component {
	return &scoreboard{
		text:         "",
		fontSize:     30,
		score:        score,
		lives:        lives,
		bonusEagleAt: 10,
	}
}

func (scoreboard scoreboard) Draw(screen *ebiten.Image) error {
	mplusNormalFont := &text.GoTextFace{
		Source: faceSource,
		Size:   float64(scoreboard.fontSize),
	}
	opWord := &text.DrawOptions{}
	opWord.GeoM.Translate(8, 4)
	opWord.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, scoreboard.text, mplusNormalFont, opWord)

	return nil
}

func (scoreboard *scoreboard) Update() error {
	scoreboard.text = fmt.Sprintf("Score: %d", *scoreboard.score)
	if *scoreboard.score == scoreboard.bonusEagleAt {
		*scoreboard.lives++
		scoreboard.bonusEagleAt += 20
	}
	return nil
}

func (scoreboard *scoreboard) OnScreen() bool {
	return true
}
