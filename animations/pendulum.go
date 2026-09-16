package animations

import "github.com/Sheriff-Hoti/gometry/shapes"

// Pendulum returns a step function that swings PointB vertically around
// the height it has on the first tick, flipping direction with clamping
// at amplitude cells above and below that base line (and at the screen
// edges on tiny screens). Values below 1 behave as 1.
func Pendulum(amplitude, cellsPerTick int) Step {
	if amplitude < 1 {
		amplitude = 1
	}
	if cellsPerTick < 1 {
		cellsPerTick = 1
	}
	dy := 1
	var baseY int
	started := false
	return func(l *shapes.Line, _, height int) {
		if l == nil || height <= 1 {
			return
		}
		if !started {
			baseY = l.PointB.Y
			started = true
		}
		l.PointB.Y += dy * cellsPerTick
		if l.PointB.Y >= baseY+amplitude || l.PointB.Y >= height-1 {
			l.PointB.Y = min(baseY+amplitude, height-1)
			dy = -1
		} else if l.PointB.Y <= baseY-amplitude || l.PointB.Y <= 0 {
			l.PointB.Y = max(baseY-amplitude, 0)
			dy = 1
		}
	}
}
