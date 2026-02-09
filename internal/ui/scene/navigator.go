package scene

import (
	"djweber/slip-hop/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Navigator struct {
	scenes     []ui.GameObject
	transition Transition
}

func (m *Navigator) Layout(w, h int) {
	for _, s := range m.scenes {
		s.Layout(w, h)
	}

	if m.transition != nil {
		m.transition.Layout(w, h)
	}
}

func (m *Navigator) Update() error {
	for _, s := range m.scenes {
		err := s.Update()
		if err != nil {
			return err
		}
	}

	if m.transition != nil {
		err := m.transition.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Navigator) Draw(screen *ebiten.Image) {
	for _, s := range m.scenes {
		s.Draw(screen)
	}
	if m.transition != nil {
		m.transition.Draw(screen)
	}
}

func (m *Navigator) NewScene() *BaseScene {
	return &BaseScene{Navigator: m}
}

//TODO refactor method to take explicit start/end behavior

// Push pushes a scene onto the navigator's scene stack
func (m *Navigator) Push(dst ui.GameObject, t Transition) {
	m.scenes = append(m.scenes, dst)

	if t != nil {
		m.transition = t
		t.Start(nil, func() {
			m.transition = nil
		})
	}
}

func NewNavigator() *Navigator {
	return &Navigator{}
}
