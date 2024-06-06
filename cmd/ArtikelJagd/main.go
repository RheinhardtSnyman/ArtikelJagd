package main

import (
	"fmt"
	"log"

	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	components []component.Component
}

const (
	north = iota
	east
	west
)

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
	game := Start()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}

func Start() *Game {

	fmt.Println("Starting")

	ebiten.SetWindowSize(800, 480)
	ebiten.SetWindowTitle("ArtikelJagd")

	game := &Game{
		components: []component.Component{
			component.NewBackground(),
			component.NewTable(),
			component.NewCurtain(east),
			component.NewCurtain(west),
			component.NewCurtain(north),
		},
	}
	// game := &Game{}
	// game.components = []component.Component{
	// 	component.NewTable(),
	// }

	return game
}

func (game *Game) Run() error {

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	return nil
}
