package stealth

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const rateFile = "rate_state.json"

type rateState struct {
	Date        string `json:"date"`
	Hour        int    `json:"hour"`
	DailyCount  int    `json:"daily_count"`
	HourlyCount int    `json:"hourly_count"`
}

func CanPerformAction(dailyLimit, hourlyLimit int) error {
	now := time.Now()
	today := now.Format("2006-01-02")
	hour := now.Hour()

	state := rateState{}
	data, _ := os.ReadFile(rateFile)
	_ = json.Unmarshal(data, &state)

	// Reset daily
	if state.Date != today {
		state.Date = today
		state.DailyCount = 0
	}

	// Reset hourly
	if state.Hour != hour {
		state.Hour = hour
		state.HourlyCount = 0
	}

	if state.DailyCount >= dailyLimit {
		return errors.New("daily limit reached")
	}
	if state.HourlyCount >= hourlyLimit {
		return errors.New("hourly limit reached")
	}

	state.DailyCount++
	state.HourlyCount++

	out, _ := json.MarshalIndent(state, "", " ")
	_ = os.WriteFile(rateFile, out, 0644)

	return nil
}
