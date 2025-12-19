package auth

import (
	"errors"
	"os"
	"time"
	"github.com/go-rod/rod"
)

func Login(page *rod.Page) error {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return errors.New("missing LINKEDIN_EMAIL or LINKEDIN_PASSWORD")
	}
	
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	page.MustElement(`input[name="session_key"]`).
		MustInput(password)
	time.Sleep(800 * time.Millisecond)

	page.MustElement(`button[type="submit"]`).
		MustClick()
	page.MustWaitLoad()

	return nil
}