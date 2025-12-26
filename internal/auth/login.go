package auth

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"
	"github.com/go-rod/rod"
	"github.com/paragraph1148/linkedin-automation/internal/config"
	"github.com/paragraph1148/linkedin-automation/internal/stealth"
)

func Login(page *rod.Page, cfg *config.Config) error {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return errors.New("missing LINKEDIN_EMAIL or LINKEDIN_PASSWORD")
	}

	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	// Email input
	emailEl := page.MustElement(`input[name="session_key"]`)
	if err := stealth.HumanAction(page, emailEl, cfg, func() error {
		stealth.HumanType(emailEl, email)
		return nil
	}); err != nil {
		return err
	}

	// Password input
	passwordEl := page.MustElement(`input[name="session_password"]`)
	if err := stealth.HumanAction(page, passwordEl, cfg, func() error {
		stealth.HumanType(passwordEl, password)
		return nil
	}); err != nil {
		return err
	}

	// Submit button
	submitBtn := page.MustElement(`button[type="submit"]`)
	if err := stealth.HumanAction(page, submitBtn, cfg, func() error {
		return stealth.HumanClick(page, submitBtn)
	}); err != nil {
		return err
	}

	page.Timeout(15 * time.Second).MustWaitLoad()

	page.MustScreenshot("after_login.png")

	url := page.MustInfo().URL
	log.Println("Post-login URL:", url)

	if strings.Contains(url, "checkpoint") || strings.Contains(url, "challenge") {
		return errors.New("linkedin security checkpoint detected")
	}

	return nil
}
