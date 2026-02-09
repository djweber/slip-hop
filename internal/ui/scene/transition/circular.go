package transition

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type CircularTransition struct {
}

func (c *CircularTransition) Layout(w, h int) {
}

func (c *CircularTransition) Update() error {
	return nil
}

func (c *CircularTransition) Draw(i *ebiten.Image) {

}

func (c *CircularTransition) Start(cb func()) {
	cb()
}
