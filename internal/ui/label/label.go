package label

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Label struct {
	Config
}

func (l *Label) Update() error {
	return nil
}

func (l *Label) Draw(screen *ebiten.Image) {
	title := l.Title
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	op.GeoM.Translate(l.X, l.Y)
	text.Draw(screen, title, l.TextFace, op)
}

type Config struct {
	Title    string
	TextFace text.Face
	X, Y     float64
}

func NewLabel(config *Config) *Label {
	return &Label{
		*config,
	}
}
