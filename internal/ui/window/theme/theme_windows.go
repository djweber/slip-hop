//go:build windows

package theme

import (
	"djweber/slip-hop/internal/ui/window/theme/windows"
	"djweber/slip-hop/internal/ui/window/theme/windows/backdrop"
)

func ApplyTheme() {
	go windows.ApplyBackdrop(backdrop.MainWindow)
}
