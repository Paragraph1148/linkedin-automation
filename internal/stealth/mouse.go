package stealth

import (
	"math/rand"
	"time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func RandomMouseMove(page *rod.Page) {
	p := proto.Point {	
		X: float64(rand.Intn(800) + 100),
		Y: float64(rand.Intn(600) + 100),
	}
	page.Mouse.MoveTo(p)
	time.Sleep(time.Duration(300+rand.Intn(700)) * time.Millisecond)
}