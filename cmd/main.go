package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/paragraph1148/linkedin-automation/internal/browser"
)

func main() {
	_ = godotenv.Load()

	b, err := browser.LaunchBrowser()
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	page := b.MustPage("https://google.com")
	page.MustWaitLoad()

	log.Println("Opened Google successfully")

	select {}
}