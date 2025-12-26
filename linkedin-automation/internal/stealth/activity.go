package stealth

import (
	"errors"
	"math/rand"
	"time"
	"github.com/paragraph1148/linkedin-automation/internal/config"
)

// EnforceSchedule blocks execution outside working hours
func EnforceSchedule(cfg *config.Config) error {
	now := time.Now()
	hour := now.Hour()

	if hour < cfg.Schedule.StartHour || hour >= cfg.Schedule.EndHour {
		return errors.New("outside scheduled activity window")
	}

	return nil
}

// RandomBreak simulates a human stepping away
func RandomBreak() {
	// 10% chance of a longer break
	if rand.Float64() < 0.10 {
		time.Sleep(time.Duration(rand.Intn(5)+3) * time.Minute)
		return
	}

	// Small pause
	time.Sleep(time.Duration(rand.Intn(1200)+600) * time.Millisecond)
}
