package stealth

import (
	"errors"
	"math/rand"
	"time"
	"github.com/go-rod/rod"
)

// HumanAction wraps a real action with human-like behavior
func HumanAction(page *rod.Page, action func() error) error {
	if page == nil {
		return errors.New("nil page")
	}

	// 1. Pre-action thinking time
	think()

	// 2. Idle mouse wandering
	idleMouse(page)

	// 3. Light scroll before action
	if rand.Float64() < 0.6 {
		RandomScroll(page)
	}

	// 4. Execute the real action
	if err := action(); err != nil {
		return err
	}

	// 5. Post-action pause (reaction time)
	react()

	return nil
}

func think() {
	time.Sleep(time.Duration(rand.Intn(1200)+400) * time.Millisecond)
}

func react() {
	time.Sleep(time.Duration(rand.Intn(800)+300) * time.Millisecond)
}

func idleMouse(page *rod.Page) {
	// Move mouse slightly without purpose
	x := rand.Float64()*300 + 200
	y := rand.Float64()*200 + 200

	page.Mouse.Move(x, y, 1)
	time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
}
