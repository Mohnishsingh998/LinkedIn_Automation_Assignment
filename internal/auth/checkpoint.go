package auth

import (
	"strings"

	"github.com/go-rod/rod"
)

func HasCheckpoint(page *rod.Page) bool {
	url := page.MustInfo().URL
	return containsAny(url, []string{
		"/checkpoint/",
		"/challenge/",
	})
}

func containsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
