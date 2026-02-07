package main

import (
	"lock-on-labs/slip-hop/internal/game"
	"lock-on-labs/slip-hop/internal/window/theme/windows"
	"lock-on-labs/slip-hop/internal/window/theme/windows/backdrop"
	"log"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title        = "SLIP-HOP"
	windowWidth  = 576
	windowHeight = 1024
)

func main() {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(windowWidth, windowHeight)

	if runtime.GOOS == "windows" {
		go windows.ApplyBackdrop(backdrop.Auto)
	}

	g := game.NewGame()

	err := ebiten.RunGame(&g)

	if err != nil {
		log.Fatal(err)
	}
}
