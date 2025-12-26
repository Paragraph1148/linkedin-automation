package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanClick(page *rod.Page, el *rod.Element) error {
	box, err := el.Box()
	if err != nil {
		return err
	}

	// Target is NOT center — humans don't aim perfectly
	target := Point{
		X: box.X + box.Width*(0.3+rand.Float64()*0.4),
		Y: box.Y + box.Height*(0.3+rand.Float64()*0.4),
	}

	// Start position (cursor "somewhere")
	start := Point{
		X: rand.Float64()*300 + 100,
		Y: rand.Float64()*300 + 100,
	}

	// Overshoot slightly
	overshoot := Point{
		X: target.X + rand.Float64()*10 - 5,
		Y: target.Y + rand.Float64()*10 - 5,
	}

	// Move → overshoot → correction
	moveBezier(page, start, overshoot)
	time.Sleep(time.Duration(rand.Intn(120)+60) * time.Millisecond)
	moveBezier(page, overshoot, target)

	// Pre-click hesitation
	time.Sleep(time.Duration(rand.Intn(150)+80) * time.Millisecond)

	page.Mouse.Click("left")

	// Post-click micro jitter (optional realism)
	if rand.Float64() < 0.3 {
		page.Mouse.Move(
			target.X+rand.Float64()*3,
			target.Y+rand.Float64()*3,
			1,
		)
	}

	return nil
}
