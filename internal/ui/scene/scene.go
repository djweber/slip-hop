package scene

import (
	"djweber/slip-hop/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

// BaseScene is a [ui.GameObject] that maintains a collection of other GameObjects.
// It includes boilerplate behavior for rendering and laying out its children.
type BaseScene struct {
	*Navigator
	children []ui.GameObject
}

func (b *BaseScene) Add(child ui.GameObject) {
	b.children = append(b.children, child)
}

func (b *BaseScene) Update() error {
	if b.children == nil {
		return nil
	}

	for _, d := range b.children {
		err := d.Update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BaseScene) Draw(img *ebiten.Image) {
	if b.children == nil {
		return
	}

	for _, d := range b.children {
		d.Draw(img)
	}
}
