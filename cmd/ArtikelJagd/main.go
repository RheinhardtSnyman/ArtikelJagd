// Entry point for application
// NOTE: main is not importable
// Can only be tested within test file in same directory which is not ideal as test directory is perfered

package main

import (
	"log"

	artikeljagd "github.com/RheinhardtSnyman/ArtikelJagd"
)

func main() {
	game := artikeljagd.Start()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
