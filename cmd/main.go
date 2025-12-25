package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/paragraph1148/linkedin-automation/internal/auth"
	"github.com/paragraph1148/linkedin-automation/internal/browser"
	"github.com/paragraph1148/linkedin-automation/internal/stealth"
	"github.com/paragraph1148/linkedin-automation/internal/search"
)

func main() {
	_ = godotenv.Load()

	useMock := os.Getenv("USE_MOCK") == "1"

	b, err := browser.LaunchBrowser()
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	page := b.MustPage()

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
		search.SaveProfiles(profiles, "profiles_mock.json")
		return
	}
	for _, c := range meta.Capabilities {
		log.Println("✔", c)
	}

	if !auth.LoadCookies(page) {
		log.Println("No cookies found, logging in")
		if err := auth.Login(page); err != nil {
			log.Fatal(err)
		}
		auth.SaveCookies(page)
	} else {
		log.Println("loaded cookies, skipping login")
	}

	page.MustNavigate("https://www.linkedin.com/feed")
	page.MustWaitLoad()

	if err := auth.EnsureAuthenticated(page); err != nil {
		log.Println("Auth guard triggered:", err)
		page.MustScreenshot("after_login.png")
		return
	}

	query := search.SearchQuery{
		Keywords: "Software Engineer",
		Page: 0,
	}
	log.Println("Navigating to search URL:", query.URL())

	page.MustNavigate(query.URL())
	page.MustWaitLoad()
	
	stealth.RandomScroll(page)
	stealth.RandomMouseMove(page)

	profiles, err := search.ExtractProfileURLs(page)
	if err != nil {
		log.Println("Profile extraction failed:", err)
		return
	}
	
	log.Println("Num of profiles found:", len(profiles))

	if err := search.SaveProfiles(profiles, "profiles.json"); err != nil {
		log.Println("Failed to save profiles:", err)
	}
}