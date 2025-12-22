package stealth

import (
	"math/rand"
	"time"
	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	scrollCount := rand.Intn(3) + 2

	for i := 0; i < scrollCount; i++ {
		page.MustEval(`
			car distance = Math.floor(Math.random() * 400) + 200;
			window.scrollBy(0, distance)`)
		time.Sleep(time.Duration(400+rand.Intn(800)) * time.Millisecond)
	}
}