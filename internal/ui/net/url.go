package net

import (
	"log"

	"github.com/pkg/browser"
)

func OpenUrl(url string) {
	err := browser.OpenURL(url)

	if err != nil {
		log.Fatal(err)
	}
}
