package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"

	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	red = iota
	blue
	green
	none
)

type Game struct {
	components []component.Component
	score      int
	armed      int
}

const (
	north = iota
	east
	west
)

func getRandomKeyValue() (int, string) {

	word := map[string]int{"red": 0, "blue": 1, "green": 2}

	keys := make([]string, 0, len(word))
	for key := range word {
		keys = append(keys, key)
	}

	index := rand.Intn(len(keys))

	key := keys[index]
	value := word[key]

	return value, key
}

func (game *Game) Update() error {
	for idx, cmpt := range game.components {
		if err := cmpt.Update(); err != nil {
			log.Fatal(err)
		}

		variety, value := getRandomKeyValue()

		if !cmpt.OnScreen() {
			game.components = append(game.components[:idx], game.components[idx+1:]...)
			game.components = slices.Insert(game.components, 2, component.NewfloatyWord(&game.score, 800, 30, &game.armed, variety, value))
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
		armed: none,
	}

	variety, value := getRandomKeyValue()

	game.components = []component.Component{
		component.NewBackground(),
		component.NewWave(true, "water2", 60, 0.4, -1, 210, 0.15, 25),
		component.NewfloatyWord(&game.score, 800, 30, &game.armed, variety, value),
		component.NewWave(false, "water1"),
		component.NewTable(),
		component.NewCurtain(east),
		component.NewCurtain(west),
		component.NewCurtain(north),

		component.NewScoreboard(&game.score),
		component.NewButton("Red", 200.00, red, &game.armed),
		component.NewButton("Blue", 350.00, blue, &game.armed),
		component.NewButton("Grn", 515.00, green, &game.armed),

		component.NewCrosshair(&game.armed),
	}

	return game
}

func (game *Game) Run() error {

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	return nil
}
