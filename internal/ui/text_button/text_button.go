package text_button

import (
	"djweber/slip-hop/internal/ui/input"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextButton struct {
	*Config
	isPressed bool
	clickable *input.Clickable
}

func (b *TextButton) Update() error {
	return b.clickable.Update()
}

func (b *TextButton) Draw(img *ebiten.Image) {
	var o int

	if b.isPressed {
		o = 1
	}

	to := &text.DrawOptions{}
	to.PrimaryAlign = text.AlignCenter   // x-axis
	to.SecondaryAlign = text.AlignCenter // y-axis
	to.GeoM.Translate(float64(b.X), float64(b.Y+o))
	to.ColorScale = ebiten.ColorScale{}
	to.Filter = ebiten.FilterNearest
	text.Draw(img, b.Text, b.TypeFace, to)
}

type Config struct {
	Text     string
	TypeFace text.Face
	X, Y     int
	OnClick  func()
}

func NewTextButton(c *Config) *TextButton {
	w, h := text.Measure(c.Text, c.TypeFace, 0)
	x0 := c.X - int(w)/2
	y0 := c.Y - int(h)/2
	x1 := x0 + int(w)
	y1 := y0 + int(h)

	b := image.Rect(x0, y0, x1, y1)

	btn := &TextButton{
		Config: c,
	}

	btn.clickable = input.NewClickable(
		b,
		16,
		c.OnClick,
		func() {
			btn.isPressed = true
		},
		func() {
			btn.isPressed = false
		},
	)

	return btn
}
