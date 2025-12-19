package browser

import (
	"log"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func LaunchBrowser() (*rod.Browser, error) {
	log.Println("Launching Rod-managed Chromium (headless, no sandbox)")

	u := launcher.New().
		Headless(true).
		NoSandbox(true).
		Leakless(false).
		MustLaunch()

	browser := rod.New().ControlURL(u)

	if err := browser.Connect(); err != nil {
		return nil, err
	}

	log.Println("Chromium launched successfully")
	return browser, nil
}