package play

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/ui/image"
	"djweber/slip-hop/internal/ui/player"
	"djweber/slip-hop/internal/ui/scene"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Play struct {
	isPaused bool
	*scene.BaseScene
	player *player.Player
}

func (p *Play) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		p.isPaused = !p.isPaused
		fmt.Printf("Paused: %t\n", p.isPaused)
	}

	if p.isPaused {
		return nil
	}

	return p.BaseScene.Update()
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
