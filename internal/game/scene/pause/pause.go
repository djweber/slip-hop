package play

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/image"
	"djweber/slip-hop/internal/ui/scene"
)

type Pause struct {
	*scene.BaseScene
	sm *scene.Navigator
}

func NewPause(m *scene.Navigator) *Pause {
	var d []ui.GameObject

	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	return &Pause{
		&scene.BaseScene{
			Children: d,
		},
		m,
	}
}
