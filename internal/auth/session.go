package auth

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFilePath = "data/cookies.json"

// SaveCookies persists browser cookies to disk
func SaveCookies(page *rod.Page) error {
	cookies := page.MustCookies()

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cookieFilePath, data, 0644)
}

// LoadCookies loads cookies from disk into the browser
func LoadCookies(page *rod.Page) bool {
	data, err := os.ReadFile(cookieFilePath)
	if err != nil {
		return false
	}

	var storedCookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return false
	}

	params := make([]*proto.NetworkCookieParam, 0, len(storedCookies))

	for _, c := range storedCookies {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}
	page.MustNavigate("https://www.linkedin.com/")
	page.MustWaitLoad()
	page.MustSetCookies(params...)
	return true
}
