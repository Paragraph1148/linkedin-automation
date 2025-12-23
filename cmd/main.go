package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/paragraph1148/linkedin-automation/internal/auth"
	"github.com/paragraph1148/linkedin-automation/internal/browser"
	"github.com/paragraph1148/linkedin-automation/internal/stealth"
	"github.com/paragraph1148/linkedin-automation/internal/search"
)

func main() {
	_ = godotenv.Load()

	b, err := browser.LaunchBrowser()
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	page := b.MustPage()
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
		return
	}

	query := search.SearchQuery{
		Keywords: "Software Engineer",
		Page: 0,
	}
	page.MustNavigate(query.URL())
	page.MustWaitLoad()
	
	stealth.RandomScroll(page)
	stealth.RandomMouseMove(page)

	
}