package auth

import(
	"encoding/json"
	"os"
	"github.com/go-rod/rod"
)

const cookieFile = "cookies.json"

func SaveCookies(page *rod.Page) error {
	cookies := page.MustCookies()
	data, _ := json.MarshalIndent(cookies, "", " ")
	return os.WriteFile(cookieFile, data, 0600)
}

func LoadCookies(page *rod.Page) bool {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return false
	}

	var cookies []*rod.cookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return false
	}

	page.MustSetCookies(cookies...)
	return true
}