package stealth

import (
	"fmt"
	"errors"
	"os"
	"strconv"
	"time"
)

const dailyLimitFile = "daily_count.txt"
const dailyLimit = 15

func CanPerformAction() error {
	today := time.Now().Format("2025-12-19")

	data, _ := os.ReadFile(dailyLimitFile)
	parts := string(data)

	var savedDate(string)
	var count int

	if parts != "" {
		fmt.Sscanf(parts, "&s &d", &savedDate &count)
	}

	if savedDate != today {
		count = 0
	}

	if count >= dailyLimit {
		return errors.New("daily action limit reached")
	}

	count++
	os.WriteFile(dailyLimitFile, []byte(today+" "+strconv.Itoa(count)), 0644)
	return nil
}