package main

import (
	"log"

	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	components []component.Component
}

func (g *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	for _, component := range game.components {
		if err := component.Draw(screen); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Layout(x, y int) (screenWidth, screenHeight int) {
	return x, y
}

func main() {
	game := NewGame()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}

func NewGame() *Game {
	ebiten.SetWindowSize(800, 480)
	ebiten.SetWindowTitle("ArtikelJagd")

	// game := &Game{
	// 	components: []component.Component{
	// 		component.NewTable(),
	// 	},
	// }
	game := &Game{}
	game.components = []component.Component{
		component.NewTable(),
	}

	return game
}

func (game *Game) Run() error {

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

	return nil
}
