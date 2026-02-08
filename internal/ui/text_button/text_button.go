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

func (b *TextButton) Draw(screen *ebiten.Image) {
	var offset int

	if b.isPressed {
		offset = 1
	}

	textOptions := &text.DrawOptions{}
	textOptions.PrimaryAlign = text.AlignCenter   // x-axis
	textOptions.SecondaryAlign = text.AlignCenter // y-axis
	textOptions.GeoM.Translate(float64(b.X+offset), float64(b.Y+offset))
	textOptions.ColorScale = ebiten.ColorScale{}
	textOptions.Filter = ebiten.FilterNearest
	text.Draw(screen, b.Text, b.TypeFace, textOptions)
}

type Config struct {
	Text     string
	TypeFace text.Face
	X, Y     int
	OnClick  func()
}

func NewTextButton(config *Config) *TextButton {
	w, h := text.Measure(config.Text, config.TypeFace, 0)
	x0 := config.X - int(w)/2
	y0 := config.Y - int(h)/2
	x1 := x0 + int(w)
	y1 := y0 + int(h)

	b := image.Rect(x0, y0, x1, y1)

	btn := &TextButton{
		Config: config,
	}

	btn.clickable = input.NewClickable(
		b,
		16,
		config.OnClick,
		func() {
			btn.isPressed = true
		},
		func() {
			btn.isPressed = false
		},
	)

	return btn
}
