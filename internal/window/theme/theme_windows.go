//go:build windows

package theme

import (
	"lock-on-labs/slip-hop/internal/window/theme/windows"
	"lock-on-labs/slip-hop/internal/window/theme/windows/backdrop"
)

func ApplyTheme() {
	go windows.ApplyBackdrop(backdrop.Auto)
}
