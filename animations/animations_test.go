package animations

import (
	"math"
	"testing"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

func TestBounce(t *testing.T) {
	t.Run("moves point B right by speed", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 10, Y: 16},
			PointB: shapes.Point{X: 30, Y: 40},
		}
		step := Bounce(2)
		step(&l, 80, 24)
		if l.PointB.X != 32 {
			t.Errorf("PointB.X = %d, want 32", l.PointB.X)
		}
		if l.PointB.Y != 40 {
			t.Errorf("PointB.Y = %d, want 40 (unchanged)", l.PointB.Y)
		}
		if l.PointA != (shapes.Point{X: 10, Y: 16}) {
			t.Errorf("PointA = %v, want unchanged {10 16}", l.PointA)
		}
	})

	t.Run("bounces off both edges", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 8, Y: 0},
		}
		step := Bounce(1)
		step(&l, 10, 24)
		if l.PointB.X != 9 {
			t.Fatalf("PointB.X = %d, want 9 (clamped to edge)", l.PointB.X)
		}
		// Walk all the way left to the other edge.
		for range 9 {
			step(&l, 10, 24)
		}
		if l.PointB.X != 0 {
			t.Fatalf("PointB.X = %d, want 0 (clamped to edge)", l.PointB.X)
		}
		step(&l, 10, 24)
		if l.PointB.X != 1 {
			t.Errorf("PointB.X = %d, want 1 (moving right again)", l.PointB.X)
		}
	})

	t.Run("speed below 1 behaves as 1", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 5, Y: 0},
		}
		Bounce(0)(&l, 80, 24)
		if l.PointB.X != 6 {
			t.Errorf("PointB.X = %d, want 6", l.PointB.X)
		}
	})
}

func TestPendulum(t *testing.T) {
	t.Run("swings around the starting height", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 10, Y: 16},
			PointB: shapes.Point{X: 30, Y: 12},
		}
		step := Pendulum(6, 1)
		for range 6 {
			step(&l, 80, 40)
		}
		if l.PointB.Y != 18 {
			t.Fatalf("PointB.Y = %d, want 18 (lower extreme)", l.PointB.Y)
		}
		for range 12 {
			step(&l, 80, 40)
		}
		if l.PointB.Y != 6 {
			t.Fatalf("PointB.Y = %d, want 6 (upper extreme)", l.PointB.Y)
		}
		if l.PointB.X != 30 {
			t.Errorf("PointB.X = %d, want 30 (unchanged)", l.PointB.X)
		}
		if l.PointA != (shapes.Point{X: 10, Y: 16}) {
			t.Errorf("PointA = %v, want unchanged {10 16}", l.PointA)
		}
	})

	t.Run("clamps to the screen on tiny heights", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 5, Y: 3},
		}
		Pendulum(6, 1)(&l, 80, 4)
		if l.PointB.Y != 3 {
			t.Errorf("PointB.Y = %d, want 3 (clamped to height-1)", l.PointB.Y)
		}
	})
}

func TestOrbit(t *testing.T) {
	t.Run("keeps the radius and returns after a full circle", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 10, Y: 0},
		}
		step := Orbit(20)
		start := l.PointB
		for i := range 20 {
			step(&l, 80, 40)
			dx := float64(l.PointB.X)
			dy := float64(l.PointB.Y)
			if got := math.Hypot(dx, dy); math.Abs(got-10) > 1 {
				t.Fatalf("step %d: radius = %v, want 10±1 (got %v)", i, got, l.PointB)
			}
		}
		if l.PointB != start {
			t.Errorf("PointB = %v, want %v after full circle", l.PointB, start)
		}
		if l.PointA != (shapes.Point{X: 0, Y: 0}) {
			t.Errorf("PointA = %v, want unchanged (pivot)", l.PointA)
		}
	})
}

func TestStepGuards(t *testing.T) {
	t.Run("nil line is a no-op", func(t *testing.T) {
		// Must not panic.
		Bounce(1)(nil, 80, 24)
		Pendulum(6, 1)(nil, 80, 24)
		Orbit(20)(nil, 80, 24)
	})

	t.Run("degenerate screens are a no-op", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 5, Y: 5},
		}
		before := l
		Bounce(1)(&l, 0, 24)
		Pendulum(6, 1)(&l, 80, 0)
		if l != before {
			t.Errorf("line = %v, want unchanged %v", l, before)
		}
	})
}
