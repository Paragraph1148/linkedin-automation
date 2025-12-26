package stealth

import (
	"math"
	"math/rand"
	"time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Point struct {
	X float64
	Y float64
}

func moveBezier(page *rod.Page, from, to Point) {
	steps := rand.Intn(20) + 25

	ctrl := Point{
		X: (from.X+to.X)/2 + rand.Float64()*80 - 40,
		Y: (from.Y+to.Y)/2 + rand.Float64()*80 - 40,
	}

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		x := math.Pow(1-t, 2)*from.X +
			2*(1-t)*t*ctrl.X +
			math.Pow(t, 2)*to.X

		y := math.Pow(1-t, 2)*from.Y +
			2*(1-t)*t*ctrl.Y +
			math.Pow(t, 2)*to.Y

		page.Mouse.MoveTo(proto.Point{X: x, Y: y})
		time.Sleep(time.Duration(rand.Intn(12)+5) * time.Millisecond)
	}
}
