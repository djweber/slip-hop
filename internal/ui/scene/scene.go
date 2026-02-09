package scene

import (
	"djweber/slip-hop/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	*Navigator
	Children []ui.Drawable
}

func (b *Scene) Layout(w, h int) {
	for _, d := range b.Children {
		d.Layout(w, h)
	}
}

func (b *Scene) Update() error {
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

func (b *Scene) Draw(img *ebiten.Image) {
	if b.Children == nil {
		return
	}

	for _, d := range b.Children {
		d.Draw(img)
	}
}

type Navigator struct {
	Current ui.Drawable
}

func (m *Navigator) Layout(w, h int) {
	m.Current.Layout(w, h)
}

func (m *Navigator) Update() error {
	return m.Current.Update()
}

func (m *Navigator) Draw(screen *ebiten.Image) {
	m.Current.Draw(screen)
}

func (m *Navigator) NewScene() *Scene {
	return &Scene{Navigator: m}
}

func (m *Navigator) Go(dst ui.Drawable, t Transition) {
	m.Current = t
	t.Start(func() {
		m.Current = dst
	})
}

func NewNavigator() *Navigator {
	return &Navigator{}
}

type Transition interface {
	ui.Drawable
	Start(cb func())
}
