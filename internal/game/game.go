package game

import (
	"djweber/slip-hop/internal/config"
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

func (g *Game) Layout(_, _ int) (_, _ int) {
	return config.LayoutWidth, config.LayoutHeight
}

type Config struct {
	Title string
}

func NewGame(config *Config) Game {
	n := scene.NewNavigator()

	m := menu.NewMenu(
		n,
		&menu.Config{
			Title: config.Title,
		},
	)

	n.Push(m, nil)

	return Game{
		navigator: n,
	}
}
