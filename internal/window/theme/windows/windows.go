package windows

import (
	"fmt"
	"lock-on-labs/slip-hop/internal/window/theme/windows/backdrop"
	"os"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32                       = syscall.NewLazyDLL("user32.dll")
	procEnumWindows              = user32.NewProc("EnumWindows")
	procGetWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
	procIsWindowVisible          = user32.NewProc("IsWindowVisible")
)

const dwmwaSystemBackdropType = 38

func findWindowByProcessId(myPid uint32) windows.HWND {
	var target windows.HWND

	cb := syscall.NewCallback(func(hwnd, lParam uintptr) uintptr {
		var pid uint32
		procGetWindowThreadProcessId.Call(hwnd, uintptr(unsafe.Pointer(&pid)))

		if myPid == pid {
			ret, _, _ := procIsWindowVisible.Call(hwnd)

			if ret != 0 {
				target = windows.HWND(hwnd)
				return 0
			}
		}
		return 1
	})

	procEnumWindows.Call(cb, 0)

	return target
}

func ApplyBackdrop(backdrop backdrop.Backdrop) {
	pid := uint32(os.Getpid())

	var window windows.HWND

	// time out if we don't fetch the window within 5 seconds
	deadline := time.Now().Add(5 * time.Second)

	for window == 0 {
		if time.Now().After(deadline) {
			return
		}
		window = findWindowByProcessId(pid)
		time.Sleep(10 * time.Millisecond)
	}

	darkMode := uint32(1) // dark mode
	if err := windows.DwmSetWindowAttribute(window,
		windows.DWMWA_USE_IMMERSIVE_DARK_MODE,
		unsafe.Pointer(&darkMode),
		uint32(unsafe.Sizeof(darkMode)),
	); err != nil {
		fmt.Printf("DWMWA_USE_IMMERSIVE_DARK_MODE failed: %v\n", err)
	}

	if err := windows.DwmSetWindowAttribute(window,
		dwmwaSystemBackdropType,
		unsafe.Pointer(&backdrop),
		uint32(unsafe.Sizeof(backdrop)),
	); err != nil {
		fmt.Printf("DWMWA_SYSTEMBACKDROP_TYPE failed: %v\n", err)
	}
}
