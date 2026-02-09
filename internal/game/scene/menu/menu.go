package menu

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/game/scene/play"
	"djweber/slip-hop/internal/ui"
	"djweber/slip-hop/internal/ui/image"
	"djweber/slip-hop/internal/ui/label"
	"djweber/slip-hop/internal/ui/scene"
	"djweber/slip-hop/internal/ui/scene/transition"
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
	manager   *scene.Manager
	drawables []ui.Drawable
}

func (m *Menu) Draw(img *ebiten.Image) {
	for _, d := range m.drawables {
		d.Draw(img)
	}
}

func (m *Menu) Update() error {
	for _, d := range m.drawables {
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

func NewMenu(sm *scene.Manager, cfg *Config) *Menu {
	var m *Menu

	var d []ui.Drawable

	// background
	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	// title
	ttf := asset.LoadTextFace(asset.Font04b03, 48)

	tl := label.NewLabel(&label.Config{
		Title:    cfg.Title,
		TextFace: ttf,
		Color:    color.White,
		X:        float64(cfg.LayoutWidth / 2),
		Y:        paddingTop,
	})

	d = append(d, tl)

	// play button
	btf := asset.LoadTextFace(asset.Font04b03, 16)

	pb := text_button.NewTextButton(&text_button.Config{
		Text:     "Play",
		TypeFace: btf,
		X:        cfg.LayoutWidth / 2,
		Y:        cfg.LayoutHeight / 2,
		OnClick: func() {
			m.manager.Load(play.NewPlay(m.manager), &transition.CircularTransition{})
		},
	})

	d = append(d, pb)

	// copyright
	ctf := asset.LoadTextFace(asset.Font04b03, 8)

	ct := fmt.Sprintf("(c) %s %d", "LOCK ON", time.Now().Year())

	ctl := label.NewLabel(&label.Config{
		Title:    ct,
		TextFace: ctf,
		Color:    color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		X:        float64(cfg.LayoutWidth / 2),
		Y:        float64(cfg.LayoutHeight - paddingBottom),
	})

	d = append(d, ctl)

	m = &Menu{
		manager:   sm,
		drawables: d,
	}

	return m
}
