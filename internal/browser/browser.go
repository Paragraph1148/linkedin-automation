package browser

import (
	"log"
	"math/rand"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/paragraph1148/linkedin-automation/internal/stealth"
)

func LaunchBrowser() (*rod.Browser, stealth.FingerprintProfile, error) {
	fp := stealth.RandomFingerprint()

	u := launcher.New().
		Headless(false).
		NoSandbox(true).
		Set("user-agent", fp.UserAgent).
		Set("window-size", 
			fmt.Sprintf("%d,%d", fp.Width, fp.Height),
		).
		MustLaunch()

	browser := rod.New().ControlURL(u)

	if err := browser.Connect(); err != nil {
		return nil, fp, err
	}

	log.Println("Chromium launched with fingerprint")
	return browser, fp, nil
}
