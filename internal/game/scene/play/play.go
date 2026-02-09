package play

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/image"
	"djweber/slip-hop/internal/ui/player"
	"djweber/slip-hop/internal/ui/scene"
)

type Play struct {
	*scene.BaseScene
}

func (p *Play) Layout(w, h int) {
	if p.Children == nil {
		var d []ui.GameObject

		bg := image.NewImage(asset.ImageBackground)

		d = append(d, bg)

		pl := player.NewPlayer(float32(w/2), float32(h))

		d = append(d, pl)

		p.Children = d
	}
}

func NewPlay(n *scene.Navigator) *Play {
	return &Play{
		&scene.BaseScene{
			Navigator: n,
		},
	}
}
