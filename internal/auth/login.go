package auth

import (
	"errors"
	"os"
	"time"
	"github.com/go-rod/rod"
	"github.com/paragraph1148/linkedin-automation/internal/stealth"
)

func Login(page *rod.Page) error {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return errors.New("missing LINKEDIN_EMAIL or LINKEDIN_PASSWORD")
	}
	
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	emailEl := page.MustElement(`input[name="session_key"]`)
	stealth.HumanType(emailEl, email)
	stealth.RandomDelay(500, 1200)

	passwordEl := page.MustElement(`input[name="session_password]`)
	stealth.HumanType(passwordEl, password)
	stealth.RandomDelay(500, 1200)
	
	page.MustElement(`button[type="submit"]`).
		MustClick()
	page.MustWaitLoad()

	return nil
}