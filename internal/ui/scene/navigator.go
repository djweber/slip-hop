package scene

import (
	"djweber/slip-hop/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Navigator struct {
	scenes     []ui.GameObject
	transition Transition
}

func (n *Navigator) Layout(w, h int) {
	for _, s := range n.scenes {
		s.Layout(w, h)
	}

	if n.transition != nil {
		n.transition.Layout(w, h)
	}
}

func (n *Navigator) Update() error {
	for _, s := range n.scenes {
		err := s.Update()
		if err != nil {
			return err
		}
	}

	if n.transition != nil {
		err := n.transition.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Navigator) Draw(screen *ebiten.Image) {
	for _, s := range n.scenes {
		s.Draw(screen)
	}
	if n.transition != nil {
		n.transition.Draw(screen)
	}
}

func (n *Navigator) NewScene() *BaseScene {
	return &BaseScene{Navigator: n}
}

//TODO refactor method to take explicit start/end behavior

// Push pushes a scene onto the navigator's scene stack
func (n *Navigator) Push(dst ui.GameObject, t Transition) {
	n.scenes = append(n.scenes, dst)

	if t != nil {
		n.transition = t
		t.Start(nil, func() {
			n.transition = nil
		})
	}
}

func (n *Navigator) Pop() {
	if len(n.scenes) == 0 {
		n.scenes = n.scenes[:len(n.scenes)-1]
	}
}

func NewNavigator() *Navigator {
	return &Navigator{}
}
