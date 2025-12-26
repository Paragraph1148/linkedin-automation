package stealth

import (
	"math/rand"
	"time"
	"github.com/paragraph1148/linkedin-automation/internal/config"
)

func RandomDelay(cfg *config.Config) {
	min := cfg.Timing.MinDelayMs
	max := cfg.Timing.MaxDelayMs
	if max <= min {
		max = min + 500
	}

	delay := rand.Intn(max-min) + min
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func ThinkTime() {
	time.Sleep(time.Duration(rand.Intn(1200)+400) * time.Millisecond)
}
