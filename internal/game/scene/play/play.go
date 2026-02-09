package play

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/image"
	"djweber/slip-hop/internal/ui/scene"
)

type Play struct {
	*scene.Scene
	sm *scene.Navigator
}

func NewPlay(m *scene.Navigator) *Play {
	var d []ui.Drawable

	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	return &Play{
		&scene.Scene{
			Children: d,
		},
		m,
	}
}
