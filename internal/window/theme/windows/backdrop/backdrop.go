package backdrop

type Backdrop uint32

// see https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwm_systembackdrop_type
const (
	Auto Backdrop = iota
	None
	MainWindow      // mica on windows 11
	TransientWindow // acrylic on windows 11
	TabbedWindow
)
