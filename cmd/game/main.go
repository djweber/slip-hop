package main

import (
	"djweber/slip-hop/internal/config"
	"djweber/slip-hop/internal/game"
	"djweber/slip-hop/internal/ui/window/theme"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title = "SLIP-HOP"
)

func main() {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(config.LayoutWidth*config.Scale, config.LayoutHeight*config.Scale)
	theme.ApplyTheme()

	cfg := &game.Config{
		Title: title,
	}

	g := game.NewGame(cfg)

	err := ebiten.RunGame(&g)

	if err != nil {
		log.Fatal(err)
	}
}
