package text_button

import (
	"image"
	"lock-on-labs/slip-hop/internal/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextButton struct {
	*Config
	clickable *input.Clickable
}

func (b *TextButton) Update() error {
	return b.clickable.Update()
}

func (b *TextButton) Draw(screen *ebiten.Image) {
	textOptions := &text.DrawOptions{}
	textOptions.PrimaryAlign = text.AlignCenter   // x-axis
	textOptions.SecondaryAlign = text.AlignCenter // y-axis
	textOptions.GeoM.Translate(float64(b.X), float64(b.Y))
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

	bounds := image.Rect(x0, y0, x1, y1)

	clickable := input.NewClickable(bounds, config.OnClick, -16)

	return &TextButton{
		config,
		clickable,
	}
}
