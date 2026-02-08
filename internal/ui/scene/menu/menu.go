package menu

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/label"
	"djweber/slip-hop/internal/ui/net"
	"djweber/slip-hop/internal/ui/text_button"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	paddingTop    = 48
	paddingBottom = 16
)

type Menu struct {
	drawables []ui.Drawable
}

func (menu *Menu) Draw(screen *ebiten.Image) {
	for _, d := range menu.drawables {
		d.Draw(screen)
	}
}

func (menu *Menu) Update() error {
	for _, d := range menu.drawables {
		err := d.Update()

		if err != nil {
			return err
		}
	}
	return nil
}

type Config struct {
	Title                     string
	LayoutWidth, LayoutHeight int
}

func NewMenu(cfg *Config) *Menu {
	drawables := make([]ui.Drawable, 0)

	titleTextFace := asset.LoadTextFace(asset.Font04b03, 48)

	titleLabel := label.NewLabel(&label.Config{
		Title:    cfg.Title,
		TextFace: titleTextFace,
		Color:    color.White,
		X:        float64(cfg.LayoutWidth / 2),
		Y:        paddingTop,
	})

	drawables = append(drawables, titleLabel)

	buttonTextFace := asset.LoadTextFace(asset.Font04b03, 16)

	playButton := text_button.NewTextButton(&text_button.Config{
		Text:     "Play",
		TypeFace: buttonTextFace,
		X:        cfg.LayoutWidth / 2,
		Y:        cfg.LayoutHeight / 2,
		OnClick: func() {
			net.OpenUrl("https://discord.gg/NjTHu6NNh4")
		},
	})

	drawables = append(drawables, playButton)

	copyrightTextFace := asset.LoadTextFace(asset.Font04b03, 8)

	copyright := fmt.Sprintf("(c) %s %d", "LOCK ON", time.Now().Year())

	copyrightLabel := label.NewLabel(&label.Config{
		Title:    copyright,
		TextFace: copyrightTextFace,
		Color:    color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x20},
		X:        float64(cfg.LayoutWidth / 2),
		Y:        float64(cfg.LayoutHeight - paddingBottom),
	})

	drawables = append(drawables, copyrightLabel)

	return &Menu{
		drawables: drawables,
	}
}
