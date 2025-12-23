package search

import (
	"strings"
	"github.com/go-rod/rod"
)

func ExtractProfileURLs(page *rod.Page) ([]string, error) {
	anchors, err := page.Elements(`a`)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{})
	var profiles []string

	for _, a := range anchors {
		href, err := a.Attribute("href")
		if err != nil || href == nil {
			continue
		}
		url := cleanURL(*href)
		if isProfileURL(url) {
			if _, ok := seen[url]; !ok {
				seen[url] = struct{}{}
				profiles = append(profiles, url)
			}
		}
	}
	return nil
}

func isProfileURL(u string) bool {
	return string.Contains(u, "/in/")
}

func cleanURL(u string) string {
	if idx := strings.Index(u, "?"); idx != -1 {
		return u[:idx]
	}
	return u
}