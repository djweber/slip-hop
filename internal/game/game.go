package game

import (
	"lock-on-labs/slip-hop/internal/scene/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ActiveScene               Scene
	RenderWidth, RenderHeight int
}

func (g *Game) Update() error {
	err := g.ActiveScene.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ActiveScene.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.RenderWidth, g.RenderHeight
}

func NewGame() Game {
	cfg := &menu.Config{}

	initialScene := menu.NewMenu(cfg)

	return Game{
		ActiveScene:  initialScene,
		RenderWidth:  238,
		RenderHeight: 512,
	}
}

type Config struct {
	InitialScene Scene
}
