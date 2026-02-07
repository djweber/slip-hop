package game

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Update() error
	Draw(screen *ebiten.Image)
}
