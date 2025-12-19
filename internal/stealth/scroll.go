package stealth

import (
	"math/rand"
	"time"
	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	scrolls := rand.Intn(3) + 2

	for i := 0; i < scrolls; i++ {
		page.MustEval(`window.scrollBy(0, Math.floor(Math.random() * 400) + 200)`)
		time.Sleep(time.Duration(400+rand.Intn(800)) * time.Millisecond)
	}
}