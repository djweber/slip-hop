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

func NewPlay(n *scene.Navigator) *Play {
	p := &Play{
		BaseScene: &scene.BaseScene{
			Navigator: n,
			Children:  nil,
		},
		player: nil,
	}

	var d []ui.GameObject

	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	pl := player.NewPlayer()
	p.player = pl

	p.Children = append(p.Children, pl)

	p.Children = d
	return p
}
