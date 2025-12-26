package messaging

import (
	"encoding/json"
	"os"
	"time"
)

const stateFile = "message_state.json"

type MessageState struct {
	ProfileURL string    `json:"profile_url"`
	Connected  bool      `json:"connected"`
	Messaged   bool      `json:"messaged"`
	LastSentAt time.Time `json:"last_sent_at"`
}

func LoadState() ([]MessageState, error) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return []MessageState{}, nil
	}

	var state []MessageState
	err = json.Unmarshal(data, &state)
	return state, err
}

func SaveState(state []MessageState) error {
	data, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}
