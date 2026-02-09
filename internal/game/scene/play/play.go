package play

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/image"
	"djweber/slip-hop/internal/ui/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct {
	sm        *scene.Manager
	drawables []ui.Drawable
}

func (p *Play) Update() error {
	for _, d := range p.drawables {
		err := d.Update()

		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Play) Draw(img *ebiten.Image) {
	for _, d := range p.drawables {
		d.Draw(img)
	}
}

func NewPlay(m *scene.Manager) *Play {
	var d []ui.Drawable

	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	return &Play{
		sm:        m,
		drawables: d,
	}
}
