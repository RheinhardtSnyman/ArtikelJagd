package app

import (
	"fmt"
	"log"
	"slices"
	"time"

	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/RheinhardtSnyman/ArtikelJagd/internal/data"
	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Name       string
	Active     bool
	components []component.Component
}

type Game struct {
	Scenes      []Scene
	Score       int
	Armed       int
	Lives       int
	LastClickAt time.Time
}

const Debouncer = 500 * time.Millisecond

const (
	defualtScore = 0
	defualtArmed = helper.NONE
	defualtLives = 3
)

func SetScene(game *Game, name string) {
	fmt.Printf("Set Scene %s\n", name)
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	for idx, scene := range game.Scenes {
		if scene.Name == name {
			game.Scenes[idx].Active = true
		} else {
			game.Scenes[idx].Active = false
		}
	}
}

// Part of game loop inicialized in Run()
// Update runs 60 times a second to give 60fps and is thus animation safe
// Runs update function of all components in sequence
func (game *Game) Update() error {
	for _, scene := range game.Scenes {
		if scene.Active {
			for idx, cmpt := range scene.components {
				if err := cmpt.Update(); err != nil {
					log.Fatal(err)
				}

				//Debouncer to prevent multiple horde of clicks
				var mouseClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && time.Now().Sub(game.LastClickAt) > Debouncer
				if mouseClicked {
					fmt.Print("Click \n")
					game.LastClickAt = time.Now()
				}

				switch scene.Name {
				case "main": //TODO multiple words should be added on screen with variance
					if game.Lives < 0 {
						SetScene(game, "end")
						break
					}
					// If floatyword component is not on screen, remove it and add new floatyword component in correct z index.
					if !cmpt.OnScreen() {
						variety, value := data.GetNoun()
						scene.components = append(scene.components[:idx], scene.components[idx+1:]...)
						scene.components = slices.Insert(scene.components, 8, component.NewfloatyWord(&game.Lives, &game.Score, 800, 30, &game.Armed, variety, value))
					}
				case "start":
					if mouseClicked {
						SetScene(game, "main")
					}
				case "end":
					if mouseClicked {
						SetScene(game, "start")
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
	for _, scene := range game.Scenes {
		if scene.Active {
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

func inicializeScenes(game *Game) {
	fmt.Printf("Inicialize Scenes \n")
	game.Lives = defualtLives
	game.Score = defualtScore
	game.Armed = defualtArmed
	// not pasing *game because game is already a pointer
	game.Scenes = []Scene{getStartScene(game, true)}
	game.Scenes = append(game.Scenes, getMainScene(game, false))
	game.Scenes = append(game.Scenes, getEndScene(game, false))
}

// Start initializes starting game components
func Start() *Game {

	fmt.Println("Starting")

	if helper.Demo {
		helper.BTN_TEXT[helper.RED] = "Red"
		helper.BTN_TEXT[helper.BLUE] = "Blue"
		helper.BTN_TEXT[helper.GREEN] = "Grn"

		helper.START_TEXT = "Start"
		helper.END_TEXT = "Game Over"
		helper.TITLE_TEXT = "Red Blue Green"
	}

	ebiten.SetWindowSize(800, 580)
	ebiten.SetWindowTitle(helper.TITLE_TEXT)

	game := &Game{}
	inicializeScenes(game)

	return game
}

func getStartScene(game *Game, active bool) Scene {
	// NOTE Comoponents are drawn stacked on each other, initialization order matters.
	return Scene{
		Name:   "start",
		Active: active,
		components: []component.Component{
			component.NewBackground("./assets/images/Stall/bg_green.png"),
			component.NewTable(),
			component.NewCurtain(helper.TOP),
			component.NewMenu(helper.START_TEXT),
		},
	}
}

func getMainScene(game *Game, active bool) Scene {
	variety, value := data.GetNoun()

	// NOTE Comoponents are drawn stacked on each other, initialization order matters.
	return Scene{
		Name:   "main",
		Active: active,
		components: []component.Component{
			component.NewBackground("./assets/images/Stall/bg_blue.png"),
			component.NewMountian(260.0, 800, 0.80),
			component.NewTree(272, 0.5, 0.18, 800, 60, 0.87),
			component.NewTree(275, 0.3, 0.3, 800, 80, 0.88),
			component.NewMountian(335.0, 800, 1),
			component.NewTree(285, 0.0, 0.45, 800, 120, 0.90),
			component.NewWave(true, "water2", 60, 0.4, -1, 210, 0.15, 25),
			component.NewfloatyWord(&game.Lives, &game.Score, 800, 30, &game.Armed, variety, value), //TODO cleanup pass game pointer once
			component.NewWave(false, "water1"),
			component.NewTable(),
			component.NewCurtain(helper.RIGHT),
			component.NewCurtain(helper.LEFT),
			component.NewCurtain(helper.TOP),

			component.NewAmmo(515, 10.0, 7, &game.Lives),
			component.NewScoreboard(&game.Lives, &game.Score), //TODO cleanup pass game pointer once
			component.NewButton(200.00, helper.RED, &game.Armed),
			component.NewButton(350.00, helper.BLUE, &game.Armed),
			component.NewButton(515.00, helper.GREEN, &game.Armed),

			component.NewCrosshair(&game.Armed),
		},
	}
}

func getEndScene(game *Game, active bool) Scene {
	// NOTE Comoponents are drawn stacked on each other, initialization order matters.
	return Scene{
		Name:   "end",
		Active: active,
		components: []component.Component{
			component.NewBackground("./assets/images/Stall/bg_red.png"),
			component.NewTable(),
			component.NewCurtain(helper.TOP),
			component.NewScoreboard(&game.Lives, &game.Score), //TODO cleanup pass game pointer once
			component.NewMenu(helper.END_TEXT),
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
