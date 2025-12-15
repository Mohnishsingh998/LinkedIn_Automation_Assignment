package auth

import (
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// IsLoggedIn checks if the LinkedIn session is authenticated
func IsLoggedIn(page *rod.Page) bool {
	// Safe SPA navigation
	_ = page.Navigate("https://www.linkedin.com/feed/")
	time.Sleep(4 * time.Second)

	url := page.MustInfo().URL

	// Hard failure cases
	if strings.Contains(url, "/login") || strings.Contains(url, "/checkpoint") {
		return false
	}

	// Signal 1: global search box (very reliable)
	if _, err := page.Timeout(5 * time.Second).
		Element(`input[placeholder*="Search"]`); err == nil {
		return true
	}

	// Signal 2: profile avatar ("Me" menu)
	if _, err := page.Timeout(5 * time.Second).
		Element(`button[aria-label*="Me"]`); err == nil {
		return true
	}

	return false
}
