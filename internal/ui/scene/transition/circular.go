package transition

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type CircularTransition struct {
}

func (c *CircularTransition) Update() error {
	return nil
}

func (c *CircularTransition) Draw(i *ebiten.Image) {

}

func (c *CircularTransition) Start(start, end func()) {
	if end != nil {
		end()
	}
}
