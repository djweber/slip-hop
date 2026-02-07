package game

import (
	"lock-on-labs/slip-hop/internal/scene/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ActiveScene               Scene
	LayoutWidth, LayoutHeight int
}

func (g *Game) Update() error {
	err := g.ActiveScene.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ActiveScene.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.LayoutWidth, g.LayoutHeight
}

type Config struct {
	Title                     string
	LayoutWidth, LayoutHeight int
}

func NewGame(config *Config) Game {
	menuConfig := &menu.Config{
		Title:        config.Title,
		LayoutWidth:  config.LayoutWidth,
		LayoutHeight: config.LayoutHeight,
	}

	return Game{
		ActiveScene:  menu.NewMenu(menuConfig),
		LayoutWidth:  config.LayoutWidth,
		LayoutHeight: config.LayoutHeight,
	}
}
