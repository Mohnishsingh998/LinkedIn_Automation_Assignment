package auth

import (
	"os"
	"time"

	"github.com/go-rod/rod"
)

func PerformLogin(page *rod.Page) error {
	_ = page.Navigate("https://www.linkedin.com/login")
	time.Sleep(2 * time.Second)

	page.MustElement(`#username`).MustInput(os.Getenv("LI_EMAIL"))
	time.Sleep(800 * time.Millisecond)

	page.MustElement(`#password`).MustInput(os.Getenv("LI_PASSWORD"))
	time.Sleep(800 * time.Millisecond)

	page.MustElement(`button[type="submit"]`).MustClick()
	time.Sleep(4 * time.Second)

	return nil
}
