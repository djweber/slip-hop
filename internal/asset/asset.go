package asset

import (
	"bytes"
	"embed"
	"image"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed font/* image/*
var fs embed.FS

const (
	ImageBackground = "image/background.png"
	Font04b03       = "font/04b03.ttf"
)

func LoadTextFace(path string, size int) text.Face {
	fontFile, err := fs.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer fontFile.Close()

	fontBytes, err := io.ReadAll(fontFile)

	if err != nil {
		log.Fatal(err)
	}

	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))

	if err != nil {
		log.Fatal(err)
	}

	face := text.GoTextFace{
		Source: fontSource,
		Size:   float64(size),
	}

	return &face
}

func LoadImage(path string) image.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(fs, path)

	if err != nil {
		log.Fatal(err)
	}

	return img
}
