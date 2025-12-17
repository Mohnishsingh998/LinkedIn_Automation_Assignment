package search

import (
	"net/url"
	"time"

	"github.com/go-rod/rod"
)

// SearchProfiles searches LinkedIn people and returns unique profile URLs
func SearchProfiles(page *rod.Page, query string, maxPages int) ([]string, error) {
	results := make(map[string]struct{})

	searchURL := buildSearchURL(query)

	// SPA-safe navigation
	_ = page.Navigate(searchURL)
	time.Sleep(4 * time.Second)

	// Soft wait for search results container (no panic)
	if _, err := page.Timeout(15 * time.Second).
		Element(`ul.reusable-search__entity-result-list`); err != nil {
		// LinkedIn sometimes delays rendering; allow fallback
		time.Sleep(3 * time.Second)
	}

	for pageNum := 1; pageNum <= maxPages; pageNum++ {
		// Human-like scroll to load results
		scrollPage(page)

		// Extract profile URLs scoped to result list
		links := page.MustElements(
			`ul.reusable-search__entity-result-list a[href*="/in/"]`,
		)

		for _, link := range links {
			href := link.MustProperty("href").String()
			results[cleanProfileURL(href)] = struct{}{}
		}

		// Pagination (safe)
		nextBtn, err := page.Element(`button[aria-label="Next"]`)
		if err != nil {
			break // no more pages
		}

		nextBtn.MustClick()
		time.Sleep(3 * time.Second)
	}

	return mapKeys(results), nil
}

// ---- Helpers ----

func buildSearchURL(query string) string {
	return "https://www.linkedin.com/search/results/people/?keywords=" +
		url.QueryEscape(query)
}

func cleanProfileURL(href string) string {
	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	u.RawQuery = ""
	return u.String()
}

func mapKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// scrollPage scrolls safely without panics or SPA issues
func scrollPage(page *rod.Page) {
	for i := 0; i < 3; i++ {
		_, _ = page.Eval(`{
			const el = document.scrollingElement || document.documentElement;
			if (el) {
				el.scrollTop = el.scrollTop + el.clientHeight;
			}
		}`)
		time.Sleep(1200 * time.Millisecond)
	}
}
