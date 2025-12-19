package auth

import(
	"encoding/json"
	"os"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
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

	var storedCookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return false
	}

	var params []*proto.NetworkCookieParam
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

	page.MustSetCookies(params...)
	return true
}