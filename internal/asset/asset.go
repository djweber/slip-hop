package asset

import (
	"bytes"
	"embed"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed font/*
var files embed.FS

const (
	Asset04b03Font = "font/04b03.ttf"
)

func LoadTextFace(path string, size int) text.Face {
	fontFile, err := files.Open(path)

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
