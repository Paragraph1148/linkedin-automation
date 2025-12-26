package search

import (
	"net/url"
	"os"
	"github.com/go-rod/rod"
)

func LoadMockHTML(page *rod.Page, path string) error {
	html, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	dataURL := "data:text/html," + url.PathEscape(string(html))
	page.MustNavigate(dataURL)
	page.MustWaitLoad()
	return nil
}