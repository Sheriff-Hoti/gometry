package animations

import (
	"math"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

// CubeStep advances a cube's rotation one animation tick.
type CubeStep func(*shapes.Cube)

// Spin returns a step function rotating a cube dx radians about X and dy
// radians about Y per tick. Angles accumulate in the closure and wrap at
// 2π, so each returned function spins its own cube independently without
// drifting over long runs.
func Spin(dx, dy float64) CubeStep {
	return func(c *shapes.Cube) {
		if c == nil {
			return
		}
		c.RotX = math.Mod(c.RotX+dx, 2*math.Pi)
		c.RotY = math.Mod(c.RotY+dy, 2*math.Pi)
	}
}
