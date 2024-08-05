// ArtikelJagd/test/app_test.go
package app

import (
	"os"
	"testing"

	app "github.com/RheinhardtSnyman/ArtikelJagd"
	component "github.com/RheinhardtSnyman/ArtikelJagd/internal/components"
	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
)

func TestSetScene(t *testing.T) {

	game := &app.Game{
		Scenes: []app.Scene{
			{Name: "start", Active: true},
			{Name: "main", Active: false},
			{Name: "end", Active: false},
		},
	}

	app.SetScene(game, "main")

	if game.Scenes[0].Active {
		t.Errorf("Expected start to be inactive")
	}
	if !game.Scenes[1].Active {
		t.Errorf("Expected main to be active")
	}
	if game.Scenes[2].Active {
		t.Errorf("Expected end to be inactive")
	}
}

func TestAddFloatyWordComponent(t *testing.T) {

	helper.DemoMode = true

	if err := os.Chdir(".."); err != nil {
		panic(err)
	}

	game := &app.Game{
		Scenes: []app.Scene{
			{Components: make([]component.Component, 20)},
		},
	}

	components := app.AddFloatyWord(game, game.Scenes[0])

	if len(components) != 21 {
		t.Errorf("Expected 1 component, got %d", len(components))
	}
}
