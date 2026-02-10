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
	player *player.Player
}

func (p *Play) Layout(w, h int) {
	if p.player == nil {
		pl := player.NewPlayer()
		p.player = pl
		p.Children = append(p.Children, pl)
	}

	// lay out our children
	p.BaseScene.Layout(w, h)
}

func (p *Play) init() {
	var d []ui.GameObject

	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	p.Children = d
}

func NewPlay(n *scene.Navigator) *Play {
	p := &Play{
		BaseScene: &scene.BaseScene{
			Navigator: n,
			Children:  nil,
		},
		player: nil,
	}

	p.init()

	return p
}
