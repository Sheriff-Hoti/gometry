package animations

import (
	"math"
	"testing"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

func TestSpin(t *testing.T) {
	t.Run("advances both angles per tick", func(t *testing.T) {
		c := shapes.Cube{RotX: 0.35, RotY: 0.5}
		Spin(0.06, 0.1)(&c)
		if math.Abs(c.RotX-0.41) > 1e-9 || math.Abs(c.RotY-0.6) > 1e-9 {
			t.Errorf("got (%v, %v), want (0.41, 0.6)", c.RotX, c.RotY)
		}
	})

	t.Run("accumulates over ticks", func(t *testing.T) {
		c := shapes.Cube{}
		step := Spin(0.1, 0.2)
		for range 5 {
			step(&c)
		}
		if math.Abs(c.RotX-0.5) > 1e-9 || math.Abs(c.RotY-1.0) > 1e-9 {
			t.Errorf("got (%v, %v), want (0.5, 1.0)", c.RotX, c.RotY)
		}
	})

	t.Run("wraps at a full turn", func(t *testing.T) {
		c := shapes.Cube{}
		Spin(2*math.Pi, 0)(&c)
		if math.Abs(c.RotX) > 1e-9 {
			t.Errorf("RotX = %v, want 0 after full turn", c.RotX)
		}
	})

	t.Run("nil cube is a no-op", func(t *testing.T) {
		// Must not panic.
		Spin(0.1, 0.1)(nil)
	})
}
