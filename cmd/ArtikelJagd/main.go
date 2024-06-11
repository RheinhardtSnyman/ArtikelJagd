package main

import (
	"fmt"
	"log"
	"slices"

	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	components []component.Component
	score      int
}

const (
	north = iota
	east
	west
)

func (game *Game) Update() error {
	for idx, cmpt := range game.components {
		if err := cmpt.Update(); err != nil {
			log.Fatal(err)
		}

		if !cmpt.OnScreen() {
			game.components = append(game.components[:idx], game.components[idx+1:]...)
			game.components = slices.Insert(game.components, 2, component.NewfloatyWord(&game.score, 800, 30))
		}
	}

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

	ebiten.SetWindowSize(800, 580)
	ebiten.SetWindowTitle("ArtikelJagd")

	game := &Game{
		score: 0,
	}

	game.components = []component.Component{
		component.NewBackground(),
		component.NewWave(true, "water2", 60, 0.4, -1, 210, 0.15, 25),
		component.NewfloatyWord(&game.score, 800, 30),
		component.NewWave(false, "water1"),
		component.NewTable(),
		component.NewCurtain(east),
		component.NewCurtain(west),
		component.NewCurtain(north),

		component.NewScoreboard(&game.score),
		component.NewButton("Red", 225.00),
		component.NewButton("Gre", 375.00),
		component.NewButton("Blu", 525.00),

		component.NewCrosshair(),
	}

	return game
}

func (game *Game) Run() error {

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	return nil
}
