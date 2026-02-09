package scene

import "djweber/slip-hop/internal/ui"

type Transition interface {
	ui.GameObject
	Start(onStart func(), onEnd func())
}
