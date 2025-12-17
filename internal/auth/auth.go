package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// IsLoggedIn checks if the LinkedIn session is authenticated
func IsLoggedIn(page *rod.Page) bool {
	// SPA-safe navigation
	_ = page.Navigate("https://www.linkedin.com/feed/")
	time.Sleep(4 * time.Second)

	url := page.MustInfo().URL

	// Hard failure cases
	if strings.Contains(url, "/login") || strings.Contains(url, "/checkpoint") {
		return false
	}

	// Signal 1: global search input (very reliable)
	if _, err := page.Timeout(5 * time.Second).
		Element(`input[placeholder*="Search"]`); err == nil {
		return true
	}

	// Signal 2: profile menu ("Me")
	if _, err := page.Timeout(5 * time.Second).
		Element(`button[aria-label*="Me"]`); err == nil {
		return true
	}

	return false
}

// Authenticate orchestrates full authentication flow
func Authenticate(page *rod.Page) error {
	// 1. Try existing session
	if LoadCookies(page) && IsLoggedIn(page) {
		return nil
	}

	// 2. Automated login
	if err := PerformLogin(page); err != nil {
		return err
	}

	// 3. Allow redirects / LinkedIn processing
	time.Sleep(5 * time.Second)

	// 4. Handle security checkpoint (2FA / captcha)
	if HasCheckpoint(page) {
		// Pause for manual completion
		time.Sleep(45 * time.Second)
	}

	// 5. Validate login success
	if !IsLoggedIn(page) {
		return errors.New("login failed")
	}

	// 6. Persist session
	return SaveCookies(page)
}
