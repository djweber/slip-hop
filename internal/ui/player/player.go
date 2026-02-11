package player

import (
	"djweber/slip-hop/internal/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	size float64 = 20
)

type Player struct {
	X, Y                     float64
	MoveSpeed                float64
	isDoubleJump             bool
	isPositioned             bool
	isGrounded               bool
	jumpHeight, jumpPeakTime float64
	vY                       float64
}

func (p *Player) Update() error {
	dt := 1.0 / 60.0
	g := 2 * p.jumpHeight / (p.jumpPeakTime * p.jumpPeakTime)

	cX := p.X
	cY := p.Y

	if !p.isGrounded {
		p.vY += g * dt
		p.Y += p.vY * dt
	}

	if !p.isGrounded && p.Y >= config.LayoutHeight-size {
		p.isGrounded = true
		p.Y = config.LayoutHeight - size
		p.isDoubleJump = false
	}

	if p.isGrounded && ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.MoveSpeed
	}

	if p.isGrounded && ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.MoveSpeed
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && p.isGrounded {
		p.jump()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !p.isGrounded && !p.isDoubleJump {
		p.isDoubleJump = true
		p.jump()
	}

	if !p.isInBounds() {
		p.X = cX
		p.Y = cY
	}

	return nil
}

func (p *Player) Draw(i *ebiten.Image) {
	vector.FillRect(i, float32(p.X), float32(p.Y), float32(size), float32(size), color.White, false)
}

func (p *Player) isInBounds() bool {
	return p.X >= 0 && p.X <= config.LayoutWidth-size && p.Y >= 0 && p.Y <= config.LayoutHeight-size
}

func (p *Player) jump() {
	p.isGrounded = false
	p.vY = -2 * p.jumpHeight / p.jumpPeakTime
}

func NewPlayer() *Player {
	return &Player{
		X:            float64(config.LayoutWidth/2) - size/2,
		Y:            float64(config.LayoutHeight) - size,
		MoveSpeed:    4.0,
		jumpHeight:   64,
		jumpPeakTime: 0.25,
		isGrounded:   true,
	}
}
