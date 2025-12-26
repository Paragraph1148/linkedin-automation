package stealth

import (
	"math/rand"
	"time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func HumanHover(page *rod.Page, el *rod.Element) error {
	shape, err := el.Shape()
	if err != nil {
		return err
	}

	box := shape.Box()

	x := box.X + box.Width*(0.2+rand.Float64()*0.6)
	y := box.Y + box.Height*(0.2+rand.Float64()*0.6)

	page.Mouse.MoveTo(proto.Point{X: x, Y: y})
	time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)

	return nil
}
