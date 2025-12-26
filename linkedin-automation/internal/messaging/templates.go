package messaging

import "fmt"

func FollowUpMessage(name string) string {
	return fmt.Sprintf(
		"Hi %s! Thanks for connecting — looking forward to exchanging ideas here.",
		name,
	)
}
