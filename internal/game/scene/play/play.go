package play

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct{}

func (p *Play) Update() error {
	return nil
}

func (p *Play) Draw(img *ebiten.Image) {
}
