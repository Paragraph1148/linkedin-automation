package stealth

import (
	"github.com/go-rod/rod"
)

// HumanClick moves mouse + clicks like a human
func HumanClick(page *rod.Page, el *rod.Element) error {
	box, err := el.Box()
	if err != nil {
		return err
	}

	x := box.X + box.Width/2
	y := box.Y + box.Height/2

	MoveMouseHuman(page, randFloat(), randFloat(), x, y)
	page.Mouse.Click("left")

	return nil
}

func randFloat() float64 {
	return 100 + (rand.Float64() * 400)
}
