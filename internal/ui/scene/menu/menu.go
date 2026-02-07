package menu

import (
	"fmt"
	"lock-on-labs/slip-hop/internal/asset"
	"lock-on-labs/slip-hop/internal/ui"
	"lock-on-labs/slip-hop/internal/ui/button"
	"lock-on-labs/slip-hop/internal/ui/label"

	"github.com/hajimehoshi/ebiten/v2"
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

	titleTextFace := asset.LoadTextFace(asset.Font04b03, 32)

	titleLabel := label.NewLabel(&label.Config{
		Title:    cfg.Title,
		TextFace: titleTextFace,
		X:        float64(cfg.LayoutWidth / 2),
		Y:        64,
	})

	drawables = append(drawables, titleLabel)

	playButton := button.NewButton(&button.Config{
		X:      cfg.LayoutWidth / 2,
		Y:      cfg.LayoutHeight / 2,
		Width:  64,
		Height: 32,
		OnClick: func() {
			fmt.Println("Clicked")
		},
	})

	drawables = append(drawables, playButton)

	return &Menu{
		drawables: drawables,
	}
}
