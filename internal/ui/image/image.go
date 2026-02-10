package image

import (
	"djweber/slip-hop/internal/asset"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image struct {
	img *ebiten.Image
}

func (i *Image) Draw(s *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(0, 0)
	s.DrawImage(i.img, o)
}

func (i *Image) Update() error {
	return nil
}

func NewImage(path string) *Image {
	img := asset.LoadImage(path)
	ei := ebiten.NewImageFromImage(img)
	return &Image{
		img: ei,
	}
}
