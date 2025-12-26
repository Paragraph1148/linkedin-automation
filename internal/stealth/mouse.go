package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func MoveMouseHuman(page *rod.Page, fromX, fromY, toX, toY float64) {
	steps := rand.Intn(15) + 20

	ctrlX := (fromX+toX)/2 + rand.Float64()*40 - 20
	ctrlY := (fromY+toY)/2 + rand.Float64()*40 - 20

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		x := math.Pow(1-t, 2)*fromX +
			2*(1-t)*t*ctrlX +
			math.Pow(t, 2)*toX

		y := math.Pow(1-t, 2)*fromY +
			2*(1-t)*t*ctrlY +
			math.Pow(t, 2)*toY

		page.Mouse.Move(x, y, 1)
		time.Sleep(time.Duration(rand.Intn(15)+5) * time.Millisecond)
	}
}
