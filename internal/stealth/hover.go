package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanHover(page *rod.Page, el *rod.Element) error {
	box, err := el.Box()
	if err != nil {
		return err
	}

	x := box.X + box.Width*(0.2+rand.Float64()*0.6)
	y := box.Y + box.Height*(0.2+rand.Float64()*0.6)

	page.Mouse.Move(x, y, 1)
	time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)

	return nil
}
