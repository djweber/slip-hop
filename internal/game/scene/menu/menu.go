package menu

import (
	"djweber/slip-hop/internal/asset"
	"djweber/slip-hop/internal/config"
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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	paddingTop    = 48
	paddingBottom = 16
)

type Menu struct {
	*Config
	*scene.BaseScene
}

func (m *Menu) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		m.start()
	}

	if err := m.BaseScene.Update(); err != nil {
		return err
	}

	return nil
}

func (m *Menu) start() {
	p := play.NewPlay(m.Navigator)
	m.Navigator.Pop()
	m.Navigator.Push(p, &transition.CircularTransition{})
}

func (m *Menu) init(title string, lw, lh int) {
	var d []ui.GameObject

	// background
	bg := image.NewImage(asset.ImageBackground)

	d = append(d, bg)

	// title
	ttf := asset.LoadTextFace(asset.Font04b03, 48)

	tl := label.NewLabel(&label.Config{
		Title:    title,
		TextFace: ttf,
		Color:    color.White,
		X:        float64(lw / 2),
		Y:        paddingTop,
	})

	d = append(d, tl)

	// play button
	btf := asset.LoadTextFace(asset.Font04b03, 16)

	pb := text_button.NewTextButton(&text_button.Config{
		Text:     "Play",
		TypeFace: btf,
		X:        lw / 2,
		Y:        lh / 2,
		OnClick: func() {
			m.start()
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
		X:        float64(lw / 2),
		Y:        float64(lh - paddingBottom),
	})

	d = append(d, ctl)

	m.Children = d
}

type Config struct {
	Title string
}

func NewMenu(sm *scene.Navigator, cfg *Config) *Menu {
	m := &Menu{
		cfg,
		sm.NewScene(),
	}

	m.init(config.Title, config.LayoutWidth, config.LayoutHeight)

	return m
}
