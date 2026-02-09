package scene

import (
	"djweber/slip-hop/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseScene struct {
	*Navigator
	Children []ui.GameObject
}

func (b *BaseScene) Layout(w, h int) {
	for _, d := range b.Children {
		d.Layout(w, h)
	}
}

func (b *BaseScene) Update() error {
	if b.Children == nil {
		return nil
	}

	for _, d := range b.Children {
		err := d.Update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BaseScene) Draw(img *ebiten.Image) {
	if b.Children == nil {
		return
	}

	for _, d := range b.Children {
		d.Draw(img)
	}
}
