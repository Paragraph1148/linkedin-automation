package stealth

import (
	"math/rand"
	"time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func HumanClick(page *rod.Page, el *rod.Element) error {
	shape, err := el.Shape()
	if err != nil {
		return err
	}

	box := shape.Box()

	targetX := box.X + box.Width*(0.3+rand.Float64()*0.4)
	targetY := box.Y + box.Height*(0.3+rand.Float64()*0.4)

	startX := rand.Float64()*300 + 100
	startY := rand.Float64()*300 + 100

	// Move mouse like a human
	page.Mouse.MoveTo(proto.Point{X: startX, Y: startY})
	time.Sleep(50 * time.Millisecond)

	page.Mouse.MoveTo(proto.Point{X: targetX, Y: targetY})
	time.Sleep(time.Duration(rand.Intn(150)+80) * time.Millisecond)

	page.Mouse.Click(proto.InputMouseButtonLeft, 1)

	return nil
}
