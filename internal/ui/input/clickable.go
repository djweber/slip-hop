package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Clickable struct {
	OnClick func()
	image.Rectangle
	padding     int
	isMouseDown bool
	isMouseUp   bool
}

func (c *Clickable) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= c.Min.X && x <= c.Max.X && y >= c.Min.Y && y <= c.Max.Y {
			c.isMouseDown = true
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
	}
	return nil
}

func (c *Clickable) Draw(*ebiten.Image) {
	// no-op
}

func NewClickable(bounds image.Rectangle, onClick func(), inset int) *Clickable {
	insetBounds := bounds.Inset(inset / 2)
	return &Clickable{
		OnClick:     onClick,
		Rectangle:   insetBounds,
		isMouseDown: false,
		isMouseUp:   true,
	}
}
