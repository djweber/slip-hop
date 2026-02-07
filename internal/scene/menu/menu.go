package menu

import (
	"image/color"
	"lock-on-labs/slip-hop/internal/asset"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Menu struct {
	layoutWidth, layoutHeight int
}

func (menu *Menu) Draw(screen *ebiten.Image) {
	title := "SLIP-HOP"
	textFace := asset.LoadTextFace(asset.Asset04b03Font, 32)
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(50, 50)
	text.Draw(screen, title, textFace, op)
}

func (menu *Menu) Update() error {
	return nil
}

type Config struct {
	LayoutWidth, LayoutHeight int
}

func NewMenu(cfg *Config) *Menu {
	return &Menu{
		layoutWidth:  cfg.LayoutWidth,
		layoutHeight: cfg.LayoutHeight,
	}
}
