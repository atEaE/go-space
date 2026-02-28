package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/atEaE/go-space/internal/game"
)

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Vampire Survivors Mini")
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
