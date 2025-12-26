package messaging

import (
	"time"
)

func RunMessagingFlow(profiles []string) error {
	state, _ := LoadState()

	accepted := MockAcceptedConnections(profiles)

	for _, profile := range accepted {
		found := false

		for i := range state {
			if state[i].ProfileURL == profile {
				found = true
				if state[i].Messaged {
					continue
				}

				msg := FollowUpMessage("there")
				_ = MockSendMessage(profile, msg)

				state[i].Messaged = true
				state[i].LastSentAt = time.Now()
			}
		}

		if !found {
			msg := FollowUpMessage("there")
			_ = MockSendMessage(profile, msg)

			state = append(state, MessageState{
				ProfileURL: profile,
				Connected:  true,
				Messaged:   true,
				LastSentAt: time.Now(),
			})
		}
	}

	return SaveState(state)
}
