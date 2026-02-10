package play

import (
	"djweber/slip-hop/internal/asset"
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
		},
		player: nil,
	}
	bg := image.NewImage(asset.ImageBackground)
	p.Add(bg)
	pl := player.NewPlayer()
	p.Add(pl)
	p.player = pl
	return p
}
