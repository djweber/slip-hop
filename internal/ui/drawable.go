package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	Update() error
	Draw(i *ebiten.Image)
	Layout(w, h int)
}
