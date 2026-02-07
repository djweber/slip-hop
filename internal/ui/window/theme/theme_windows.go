//go:build windows

package theme

import (
	"lock-on-labs/slip-hop/internal/ui/window/theme/windows"
	"lock-on-labs/slip-hop/internal/ui/window/theme/windows/backdrop"
)

func ApplyTheme() {
	go windows.ApplyBackdrop(backdrop.MainWindow)
}
