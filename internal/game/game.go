package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	RenderWidth, RenderHeight int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.RenderWidth, g.RenderHeight
}

func NewGame() Game {
	return Game{
		RenderWidth:  288,
		RenderHeight: 512,
	}
}
