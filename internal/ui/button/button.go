package button

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	*Config
	image.Rectangle
	isMouseDown bool
	isMouseUp   bool
}

func (b *Button) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= b.Min.X && x <= b.Max.X && y >= b.Min.Y && y <= b.Max.Y {
			b.isMouseDown = true
		} else {
			b.isMouseDown = false
		}
	} else {
		if b.isMouseDown {
			if b.OnClick != nil {
				b.OnClick()
			}
		}
		b.isMouseDown = false
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

	return &Button{
		config,
		bounds,
		false,
		true,
	}
}
