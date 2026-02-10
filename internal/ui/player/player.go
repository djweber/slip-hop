package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	size float64 = 32
)

type Player struct {
	X, Y                     float64
	MoveSpeed                float64
	isDoubleJump             bool
	isPositioned             bool
	isGrounded               bool
	lw, lh                   float64
	jumpHeight, jumpPeakTime float64
	vY                       float64
}

func (p *Player) Update() error {
	dt := 1.0 / 60.0
	g := 2 * p.jumpHeight / (p.jumpPeakTime * p.jumpPeakTime)

	if !p.isGrounded {
		p.vY += g * dt
		p.Y += p.vY * dt
	}

	if !p.isGrounded && p.Y >= p.lh-size {
		p.isGrounded = true
		p.Y = p.lh - size
		p.isDoubleJump = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.MoveSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.MoveSpeed
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && p.isGrounded {
		p.jump()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !p.isGrounded && !p.isDoubleJump {
		p.isDoubleJump = true
		p.jump()
	}

	return nil
}

func (p *Player) jump() {
	p.isGrounded = false
	p.vY = -2 * p.jumpHeight / p.jumpPeakTime
}

func (p *Player) Draw(i *ebiten.Image) {
	vector.FillRect(i, float32(p.X), float32(p.Y), float32(size), float32(size), color.White, false)
}

func (p *Player) Layout(w, h int) {
	p.lw = float64(w)
	p.lh = float64(h)

	if !p.isPositioned {
		p.X = p.lw/2 - size/2
		p.Y = p.lh - size
		p.isPositioned = true
	}
}

func NewPlayer() *Player {
	return &Player{
		jumpHeight:   100,
		jumpPeakTime: 0.25,
		isGrounded:   true,
		MoveSpeed:    6.0,
	}
}
