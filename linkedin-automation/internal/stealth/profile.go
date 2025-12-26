package stealth

import (
	"math/rand"
)

type FingerprintProfile struct {
	UserAgent string
	Width     int
	Height    int
	Languages []string
	Platform  string
	HardwareC int
	Plugins   int
	WebGLVend string
	WebGLRend string
}

func RandomFingerprint() FingerprintProfile {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
	}

	width := rand.Intn(400) + 1200
	height := rand.Intn(300) + 700

	return FingerprintProfile{
		UserAgent: userAgents[rand.Intn(len(userAgents))],
		Width:     width,
		Height:    height,
		Languages: []string{"en-US", "en"},
		Platform:  "Win32",
		HardwareC: rand.Intn(4) + 4, // 4–8 cores
		Plugins:   rand.Intn(3) + 3,
		WebGLVend: "Intel Inc.",
		WebGLRend: "Intel Iris OpenGL Engine",
	}
}
