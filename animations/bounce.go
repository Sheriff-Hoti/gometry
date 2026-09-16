// Package animations provides one step function per line animation.
// Each animation is a factory (Bounce, Pendulum, Orbit) returning a Step
// that advances a shapes.Line one tick. The direction, phase, and
// geometry snapshots live inside the returned closure, so every Step
// animates its own line independently and the callers stay stateless.
package animations

import "github.com/Sheriff-Hoti/gometry/shapes"

// Step advances a line one animation tick. Width and height are the
// screen size in cells, used for edge detection and clamping.
type Step func(l *shapes.Line, width, height int)

// Bounce returns a step function that moves PointB horizontally
// cellsPerTick cells per tick, bouncing off the left and right screen
// edges with clamping. Values below 1 behave as 1.
func Bounce(cellsPerTick int) Step {
	if cellsPerTick < 1 {
		cellsPerTick = 1
	}
	dx := 1
	return func(l *shapes.Line, width, _ int) {
		if l == nil || width <= 1 {
			return
		}
		l.PointB.X += dx * cellsPerTick
		if l.PointB.X >= width-1 {
			l.PointB.X = width - 1
			dx = -1
		} else if l.PointB.X <= 0 {
			l.PointB.X = 0
			dx = 1
		}
	}
}
