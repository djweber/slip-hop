package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Clickable struct {
	OnClick     func()
	OnMouseDown func()
	OnMouseUp   func()
	image.Rectangle
	padding     int
	isMouseDown bool
}

func (c *Clickable) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= c.Min.X && x <= c.Max.X && y >= c.Min.Y && y <= c.Max.Y {
			c.isMouseDown = true
			if c.OnMouseDown != nil {
				c.OnMouseDown()
			}
		} else {
			c.isMouseDown = false
		}
	} else {
		if c.isMouseDown {
			if c.OnClick != nil {
				c.OnClick()
			}
		}

		c.isMouseDown = false

		if c.OnMouseUp != nil {
			c.OnMouseUp()
		}
	}
	return nil
}

func (c *Clickable) Draw(*ebiten.Image) {
	// no-op
}

// NewClickable creates a new Clickable to attach to other Children. The target bounds b
// can be extended by providing a margin m.
func NewClickable(b image.Rectangle, m int, oc func(), omd func(), omu func()) *Clickable {
	insetBounds := b.Inset(-m / 2)
	return &Clickable{
		OnClick:     oc,
		OnMouseDown: omd,
		OnMouseUp:   omu,
		Rectangle:   insetBounds,
		isMouseDown: false,
	}
}
