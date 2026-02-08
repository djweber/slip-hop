package game

import (
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/scene/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	activeScene               ui.Scene
	layoutWidth, layoutHeight int
}

func (g *Game) Update() error {
	err := g.activeScene.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.activeScene.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.layoutWidth, g.layoutHeight
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
		activeScene:  menu.NewMenu(menuConfig),
		layoutWidth:  config.LayoutWidth,
		layoutHeight: config.LayoutHeight,
	}
}
