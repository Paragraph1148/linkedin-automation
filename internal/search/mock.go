package search

import (
	"os"
	"github.com/go-rod/rod"
)

func LoadMockHTML(page *rod.Page, path string) error {
	html, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	page.MustSetContent(string(html))
	return nil
}