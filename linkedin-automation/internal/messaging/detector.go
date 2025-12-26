package messaging

import "math/rand"

// MockAcceptedConnections simulates accepted LinkedIn connections
func MockAcceptedConnections(profiles []string) []string {
	var accepted []string

	for _, p := range profiles {
		// ~30% acceptance rate
		if rand.Float64() < 0.3 {
			accepted = append(accepted, p)
		}
	}

	return accepted
}
