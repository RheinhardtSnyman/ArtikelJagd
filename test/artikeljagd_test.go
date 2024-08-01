// ArtikelJagd/test/app_test.go
package app

import (
	"testing"

	app "github.com/RheinhardtSnyman/ArtikelJagd"
	"github.com/RheinhardtSnyman/ArtikelJagd/internal/helper"
)

func TestSetScene(t *testing.T) {

	helper.Demo = true

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
