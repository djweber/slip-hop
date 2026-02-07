package button

import (
	"image"
	"image/color"
	"lock-on-labs/slip-hop/internal/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	*Config
	*image.Rectangle
	clickable *input.Clickable
}

func (b *Button) Update() error {
	err := b.clickable.Update()
	if err != nil {
		return err
	}
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(b.Min.X), float32(b.Min.Y), float32(b.Dx()), float32(b.Dy()), color.White, true)
}

type Config struct {
	X, Y, Width, Height int
	OnClick             func()
}

func NewButton(config *Config) *Button {
	bx := config.X - config.Width/2

	by := config.Y - config.Height/2

	bounds := image.Rect(bx, by, bx+config.Width, by+config.Height)

	clickable := input.NewClickable(bounds, config.OnClick)

	return &Button{
		config,
		&bounds,
		clickable,
	}
}
