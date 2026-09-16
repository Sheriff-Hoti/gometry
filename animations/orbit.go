package animations

import (
	"math"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

// Orbit returns a step function that rotates PointB around PointA,
// completing a full circle every stepsPerCircle ticks. The pivot, radius,
// and starting phase are snapshotted from the line on the first tick.
// PointB is reprojected from polar coordinates every tick, so rounding
// error cannot accumulate and drift the radius. Values below 1 behave
// as 1.
func Orbit(stepsPerCircle int) Step {
	if stepsPerCircle < 1 {
		stepsPerCircle = 1
	}
	advance := 2 * math.Pi / float64(stepsPerCircle)
	var angle, radius float64
	var pivot shapes.Point
	started := false
	return func(l *shapes.Line, _, _ int) {
		if l == nil {
			return
		}
		if !started {
			pivot = l.PointA
			dx := float64(l.PointB.X - l.PointA.X)
			dy := float64(l.PointB.Y - l.PointA.Y)
			radius = math.Hypot(dx, dy)
			if radius < 1 {
				radius = 1
			}
			angle = math.Atan2(dy, dx)
			started = true
		}
		angle += advance
		l.PointB.X = pivot.X + int(math.Round(radius*math.Cos(angle)))
		l.PointB.Y = pivot.Y + int(math.Round(radius*math.Sin(angle)))
	}
}
