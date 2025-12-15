package browser

import "github.com/go-rod/rod"

func ApplyStealth(page *rod.Page) {
	page.MustEval(`
		Object.defineProperty(navigator, 'webdriver', { get: () => undefined });
		Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'] });
		Object.defineProperty(navigator, 'plugins', { get: () => [1, 2, 3] });
	`)
}
