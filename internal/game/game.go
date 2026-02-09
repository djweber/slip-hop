package game

import (
	"djweber/slip-hop/internal/game/scene/menu"
	"djweber/slip-hop/internal/ui/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	navigator                 *scene.Navigator
	layoutWidth, layoutHeight int
}

func (g *Game) Update() error {
	err := g.navigator.Update()
	return err
}

func (g *Game) Draw(img *ebiten.Image) {
	g.navigator.Draw(img)
}

func (g *Game) Layout(_, _ int) (lw, lh int) {
	g.navigator.Layout(g.layoutWidth, g.layoutHeight)
	return g.layoutWidth, g.layoutHeight
}

type Config struct {
	Title                     string
	LayoutWidth, LayoutHeight int
}

func NewGame(config *Config) Game {
	n := scene.NewNavigator()

	m := menu.NewMenu(
		n,
		&menu.Config{
			Title: config.Title,
		},
	)

	n.Current = m

	return Game{
		navigator:    n,
		layoutWidth:  config.LayoutWidth,
		layoutHeight: config.LayoutHeight,
	}
}
