package main

import (
	"lock-on-labs/slip-hop/internal/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title        = "SLIP HOP"
	windowWidth  = 576
	windowHeight = 1024
)

func main() {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(windowWidth, windowHeight)

	g := game.NewGame()

	err := ebiten.RunGame(&g)

	if err != nil {
		log.Fatal(err)
	}
}
