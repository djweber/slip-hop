package scene

import (
	"djweber/slip-hop/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	ui.Drawable
}

type Manager struct {
	Current Scene
}

func (m *Manager) Update() error {
	return m.Current.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.Current.Draw(screen)
}

func (m *Manager) Load(s Scene, t Transition) {
	m.Current = t

	t.Start(func() {
		m.Current = s
	})
}

func NewManager() *Manager {
	return &Manager{}
}

type Transition interface {
	Scene
	Start(cb func())
}
