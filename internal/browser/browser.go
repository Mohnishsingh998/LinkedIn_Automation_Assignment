package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func NewBrowser() *rod.Browser {
	u := launcher.New().
		Bin("/Applications/Brave Browser.app/Contents/MacOS/Brave Browser").
		Headless(false).
		MustLaunch()

	return rod.New().ControlURL(u).MustConnect()
}
