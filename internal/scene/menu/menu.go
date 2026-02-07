package menu

import (
	"image/color"
	"lock-on-labs/slip-hop/internal/asset"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Menu struct {
	title                     string
	textFace                  text.Face
	layoutWidth, layoutHeight int
}

func (menu *Menu) Draw(screen *ebiten.Image) {
	title := menu.title
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	op.GeoM.Translate(float64(menu.layoutWidth/2), 64)
	text.Draw(screen, title, menu.textFace, op)
}

func (menu *Menu) Update() error {
	return nil
}

type Config struct {
	Title                     string
	LayoutWidth, LayoutHeight int
}

func NewMenu(cfg *Config) *Menu {
	textFace := asset.LoadTextFace(asset.Font04b03, 32)

	return &Menu{
		title:        cfg.Title,
		textFace:     textFace,
		layoutWidth:  cfg.LayoutWidth,
		layoutHeight: cfg.LayoutHeight,
	}
}
