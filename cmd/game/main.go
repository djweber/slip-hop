package main

import (
	"lock-on-labs/slip-hop/internal/game"
	"lock-on-labs/slip-hop/internal/window/theme"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title        = "SLIP-HOP"
	windowWidth  = 576
	windowHeight = 1024
	layoutWidth  = windowWidth / 2
	layoutHeight = windowHeight / 2
)

func main() {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(windowWidth, windowHeight)

	theme.ApplyTheme()

	config := &game.Config{
		Title:        title,
		LayoutWidth:  layoutWidth,
		LayoutHeight: layoutHeight,
	}

	g := game.NewGame(config)

	err := ebiten.RunGame(&g)

	if err != nil {
		log.Fatal(err)
	}
}
