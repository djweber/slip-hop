package label

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Label struct {
	Config
}

func (l *Label) Layout(w, h int) {
}

func (l *Label) Update() error {
	return nil
}

func (l *Label) Draw(img *ebiten.Image) {
	do := &text.DrawOptions{}
	do.ColorScale.ScaleWithColor(l.Color)
	do.PrimaryAlign = text.AlignCenter
	do.SecondaryAlign = text.AlignCenter
	do.GeoM.Translate(l.X, l.Y)
	do.Filter = ebiten.FilterNearest
	text.Draw(img, l.Title, l.TextFace, do)
}

type Config struct {
	Title    string
	TextFace text.Face
	Color    color.Color
	X, Y     float64
}

func NewLabel(c *Config) *Label {
	return &Label{
		*c,
	}
}
