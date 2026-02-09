package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	size float32 = 32
)

type Player struct {
	w, h int
	x, y float32
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.x -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x += 1
	}

	return nil
}

func (p *Player) Draw(i *ebiten.Image) {
	vector.FillRect(i, p.x, p.y, size, size, color.White, false)
}

func (p *Player) Layout(w, h int) {
	// no-op
}

func NewPlayer(x, y float32) *Player {
	return &Player{x: x - size/2, y: y - size}
}
