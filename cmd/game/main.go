package main

import (
	"djweber/slip-hop/internal/game"
	"djweber/slip-hop/internal/ui/window/theme"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title        = "SLIP-HOP"
	layoutWidth  = 288
	layoutHeight = 512
	scale        = 2
)

func main() {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(layoutWidth*scale, layoutHeight*scale)
	theme.ApplyTheme()

	cfg := &game.Config{
		Title:        title,
		LayoutWidth:  layoutWidth,
		LayoutHeight: layoutHeight,
	}

	g := game.NewGame(cfg)

	err := ebiten.RunGame(&g)

	if err != nil {
		log.Fatal(err)
	}
}
