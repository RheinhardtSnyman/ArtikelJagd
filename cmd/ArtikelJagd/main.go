// Entry point for application
// NOTE: main is not importable
// Can only be tested with test file in same directory which is not ideal as test directory is perfered

package main

import (
	"flag"
	"log"

	app "github.com/RheinhardtSnyman/ArtikelJagd"
)

var flagDemoMode = flag.Bool("demo", false, "Demo mode")

func main() {

	flag.Parse()

	game := app.Start(flagDemoMode)
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
