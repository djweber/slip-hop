package game

import (
	"djweber/slip-hop/internal/game/scene/menu"
	"djweber/slip-hop/internal/ui/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager              *scene.Manager
	layoutWidth, layoutHeight int
}

func (g *Game) Update() error {
	err := g.sceneManager.Update()
	return err
}

func (g *Game) Draw(img *ebiten.Image) {
	g.sceneManager.Draw(img)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.layoutWidth, g.layoutHeight
}

type Config struct {
	Title                     string
	LayoutWidth, LayoutHeight int
}

func NewGame(config *Config) Game {
	sm := scene.NewManager()

	m := menu.NewMenu(
		sm,
		&menu.Config{
			Title:        config.Title,
			LayoutWidth:  config.LayoutWidth,
			LayoutHeight: config.LayoutHeight,
		},
	)

	sm.Current = m

	return Game{
		sceneManager: sm,
		layoutWidth:  config.LayoutWidth,
		layoutHeight: config.LayoutHeight,
	}
}
