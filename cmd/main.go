package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/paragraph1148/linkedin-automation/internal/auth"
	"github.com/paragraph1148/linkedin-automation/internal/browser"
	"github.com/paragraph1148/linkedin-automation/internal/config"
	"github.com/paragraph1148/linkedin-automation/internal/meta"
	"github.com/paragraph1148/linkedin-automation/internal/search"
	"github.com/paragraph1148/linkedin-automation/internal/stealth"
)

func main() {
	_ = godotenv.Load()

	useMock := os.Getenv("USE_MOCK") == "1"

	// Load config
	cfg := config.Load("config.yaml")

	// Launch browser with fingerprint
	browserInstance, fp, err := browser.LaunchBrowser()
	if err != nil {
		log.Fatal(err)
	}
	defer browserInstance.Close()

	page := browserInstance.MustPage()

	// Apply fingerprint JS overrides
	if err := stealth.ApplyFingerprint(page, fp); err != nil {
		log.Println("fingerprint apply failed:", err)
	}

	// ---- MOCK MODE ----
	if useMock {
		log.Println("Running in mock mode")

		if err := search.LoadMockHTML(page, "testdata/search_mock.html"); err != nil {
			log.Fatal(err)
		}

		profiles, err := search.ExtractProfileURLs(page)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Mock profiles found:", len(profiles))
		_ = search.SaveProfiles(profiles, "profiles_mock.json")
		return
	}

	// Print capabilities (demo-friendly)
	for _, c := range meta.Capabilities {
		log.Println("✔", c)
	}

	// Load cookies or login
	if !auth.LoadCookies(page) {
		log.Println("No cookies found, logging in")
		if err := auth.Login(page, cfg); err != nil {
			log.Fatal(err)
		}
		auth.SaveCookies(page)
	} else {
		log.Println("Loaded cookies, skipping login")
	}

	// Verify authentication
	page.MustNavigate("https://www.linkedin.com/feed")
	page.MustWaitLoad()

	if err := auth.EnsureAuthenticated(page); err != nil {
		log.Println("Auth guard triggered:", err)
		page.MustScreenshot("after_login.png")
		return
	}

	// ---- SEARCH ----
	query := search.SearchQuery{
		Keywords: "Software Engineer",
		Page:     0,
	}

	log.Println("Navigating to search:", query.URL())

	if err := stealth.HumanAction(page, nil, cfg, func() error {
		page.MustNavigate(query.URL())
		page.MustWaitLoad()
		return nil
	}); err != nil {
		log.Println("Navigation blocked:", err)
		return
	}

	profiles, err := search.ExtractProfileURLs(page)
	if err != nil {
		log.Println("Profile extraction failed:", err)
		return
	}

	log.Println("Profiles found:", len(profiles))

	if err := search.SaveProfiles(profiles, "profiles.json"); err != nil {
		log.Println("Failed to save profiles:", err)
	}
}
