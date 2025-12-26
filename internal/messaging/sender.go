package messaging

import (
	"log"
	"time"
)

// MockSendMessage simulates sending a LinkedIn message
func MockSendMessage(profileURL, message string) error {
	log.Printf("📨 Sending message to %s\n", profileURL)
	log.Printf("💬 Message: %q\n", message)

	time.Sleep(800 * time.Millisecond)
	return nil
}
