package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/paragraph1148/linkedin-automation/internal/config"
)

// HumanAction wraps any browser interaction in full human behavior simulation
func HumanAction(
	page *rod.Page,
	el *rod.Element,
	cfg *config.Config,
	action func() error,
) error {

	// Enforce work hours
	if err := EnforceSchedule(cfg); err != nil {
		return err
	}

	// Thinking
	time.Sleep(time.Duration(rand.Intn(1200)+400) * time.Millisecond)

	// Hover intent
	if el != nil {
		_ = HumanHover(page, el)
	}

	// Optional scroll
	if rand.Float64() < 0.5 {
		RandomScroll(page)
	}

	// Rate limit (example: 10 daily, 3 hourly)
	if err := CanPerformAction(cfg.Limits.DailyConnections, 3); err != nil {
		return err
	}

	// Execute real action
	if err := action(); err != nil {
		return err
	}

	// Reaction time
	time.Sleep(time.Duration(rand.Intn(900)+300) * time.Millisecond)

	// Break simulation
	RandomBreak()

	return nil
}
