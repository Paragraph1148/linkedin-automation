package stealth

import (
	"math/rand"
	"time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func RandomScroll(page *rod.Page) {
	scrolls := rand.Intn(3) + 2

	for i := 0; i < scrolls; i++ {
		page.Keyboard.Press(input.PageDown)

		time.Sleep(time.Duration(400+rand.Intn(800)) * time.Millisecond)
	}
}
