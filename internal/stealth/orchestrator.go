package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HumanAction wraps a real browser action with realistic human behavior
func HumanAction(page *rod.Page, el *rod.Element, action func() error) error {
	// 1. Thinking / hesitation before acting
	think()

	// 2. Hover target if present
	if el != nil {
		_ = HumanHover(page, el)
	}

	// 3. Optional scroll before action
	if rand.Float64() < 0.5 {
		RandomScroll(page)
	}

	// 4. Execute real action
	if err := action(); err != nil {
		return err
	}

	// 5. Reaction time after action
	react()

	return nil
}

func think() {
	time.Sleep(time.Duration(rand.Intn(1200)+400) * time.Millisecond)
}

func react() {
	time.Sleep(time.Duration(rand.Intn(900)+300) * time.Millisecond)
}
