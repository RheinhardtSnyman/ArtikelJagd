package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"time"

	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	name       string
	active     bool
	components []component.Component
}

type Game struct {
	scenes      []Scene
	score       int
	armed       int
	lives       int
	lastClickAt time.Time
}

const debouncer = 200 * time.Millisecond

const (
	defualtScore = 0
	defualtArmed = helper.NONE
	defualtLives = 3
)

func getRandomKeyValue() (int, string) {

	word := map[string]int{"red": helper.RED, "blue": helper.BLUE, "green": helper.GREEN} //TODO replace with nouns matching gender articles Die Der Das

	keys := make([]string, 0, len(word))
	for key := range word {
		keys = append(keys, key)
	}

	index := rand.Intn(len(keys))

	key := keys[index]
	value := word[key]

	return value, key
}

func setScene(game *Game, name string) bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && time.Now().Sub(game.lastClickAt) > debouncer { //debouncer to prevent multiple horde of clicks
		fmt.Printf("Set Scene %s\n", name)
		game.lastClickAt = time.Now()
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
		for idx, scene := range game.scenes {
			if scene.name == name {
				game.scenes[idx].active = true
			} else {
				game.scenes[idx].active = false
			}

		}
		return true
	}
	return false
}

// Part of game loop inicialized in Run()
// Update runs 60 times a second to give 60fps and is thus animation safe
// Runs update function of all components in sequence
func (game *Game) Update() error {
	for _, scene := range game.scenes {
		if scene.active {
			for idx, cmpt := range scene.components {
				if err := cmpt.Update(); err != nil {
					log.Fatal(err)
				}
				switch scene.name {
				case "main":
					if game.lives < 0 {
						setScene(game, "end")
						break
					}
					variety, value := getRandomKeyValue()
					// If floatyword component is not on screen, remove it and add new floatyword component in correct z index.
					if !cmpt.OnScreen() {
						scene.components = append(scene.components[:idx], scene.components[idx+1:]...)
						scene.components = slices.Insert(scene.components, 8, component.NewfloatyWord(&game.lives, &game.score, 800, 30, &game.armed, variety, value))
					}
				case "start":
					setScene(game, "main")
				case "end":
					if setScene(game, "start") {
						inicializeScenes(game)
					}
				}
			}
		}
	}
	return nil
}

// Part of game loop inicialized in Run()
// Draw runs x times a second depending on PC spec and is thus too inconsistent arcross devices to use as animation tick.
// Runs draw function of all components in sequence
func (game *Game) Draw(screen *ebiten.Image) {
	for _, scene := range game.scenes {
		if scene.active {
			for _, component := range scene.components {
				if err := component.Draw(screen); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func (g *Game) Layout(x, y int) (screenWidth, screenHeight int) {
	return x, y
}

// Inicial execution function
func main() {
	game := Start()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}

func inicializeScenes(game *Game) {
	fmt.Printf("Inicialize scenes \n")
	game.lives = defualtLives
	game.score = defualtScore
	game.armed = defualtArmed
	// not pasing *game because game is already a pointer
	game.scenes = []Scene{getStartScene(game, true)}
	game.scenes = append(game.scenes, getMainScene(game, false))
	game.scenes = append(game.scenes, getEndScene(game, false))
}

// Start initializes starting game components
func Start() *Game {

	fmt.Println("Starting")

	ebiten.SetWindowSize(800, 580)
	ebiten.SetWindowTitle("ArtikelJagd")

	game := &Game{}
	inicializeScenes(game)

	return game
}

func getStartScene(game *Game, active bool) Scene {
	// NOTE Comoponents are drawn stacked on each other, initialization order matters.
	return Scene{
		name:   "start",
		active: active,
		components: []component.Component{
			component.NewBackground("./assets/images/Stall/bg_green.png"),
			component.NewTable(),
			component.NewCurtain(helper.TOP),
			component.NewMenu("Start"),
		},
	}
}

func getMainScene(game *Game, active bool) Scene {
	variety, value := getRandomKeyValue()

	// NOTE Comoponents are drawn stacked on each other, initialization order matters.
	return Scene{
		name:   "main",
		active: active,
		components: []component.Component{
			component.NewBackground("./assets/images/Stall/bg_blue.png"),
			component.NewMountian(260.0, 800, 0.80),
			component.NewTree(272, 0.5, 0.18, 800, 60, 0.87),
			component.NewTree(275, 0.3, 0.3, 800, 80, 0.88),
			component.NewMountian(335.0, 800, 1),
			component.NewTree(285, 0.0, 0.45, 800, 120, 0.90),
			component.NewWave(true, "water2", 60, 0.4, -1, 210, 0.15, 25),
			component.NewfloatyWord(&game.lives, &game.score, 800, 30, &game.armed, variety, value), //TODO cleanup pass game pointer once
			component.NewWave(false, "water1"),
			component.NewTable(),
			component.NewCurtain(helper.RIGHT),
			component.NewCurtain(helper.LEFT),
			component.NewCurtain(helper.TOP),

			component.NewAmmo(515, 10.0, 7, &game.lives),
			component.NewScoreboard(&game.lives, &game.score), //TODO cleanup pass game pointer once
			component.NewButton("Red", 200.00, helper.RED, &game.armed),
			component.NewButton("Blue", 350.00, helper.BLUE, &game.armed),
			component.NewButton("Grn", 515.00, helper.GREEN, &game.armed),

			component.NewCrosshair(&game.armed),
		},
	}
}

func getEndScene(game *Game, active bool) Scene {
	// NOTE Comoponents are drawn stacked on each other, initialization order matters.
	return Scene{
		name:   "end",
		active: active,
		components: []component.Component{
			component.NewBackground("./assets/images/Stall/bg_red.png"),
			component.NewTable(),
			component.NewCurtain(helper.TOP),
			component.NewScoreboard(&game.lives, &game.score), //TODO cleanup pass game pointer once
			component.NewMenu("Game Over"),
		},
	}
}

// Executes game loop
func (game *Game) Run() error {

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	return nil
}
