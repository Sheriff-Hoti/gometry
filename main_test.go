package main

import (
	"testing"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

func TestAnimate(t *testing.T) {
	t.Run("moves point B one cell right", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 10, Y: 16},
			PointB: shapes.Point{X: 30, Y: 40},
		}
		animate(&l)
		if l.PointB.X != 31 {
			t.Errorf("PointB.X = %d, want 31", l.PointB.X)
		}
		if l.PointB.Y != 40 {
			t.Errorf("PointB.Y = %d, want 40 (unchanged)", l.PointB.Y)
		}
		if l.PointA != (shapes.Point{X: 10, Y: 16}) {
			t.Errorf("PointA = %v, want unchanged {10 16}", l.PointA)
		}
	})

	t.Run("accumulates over ticks", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 5, Y: 5},
		}
		for range 3 {
			animate(&l)
		}
		if l.PointB.X != 8 {
			t.Errorf("PointB.X = %d, want 8", l.PointB.X)
		}
	})

	t.Run("nil line is a no-op", func(t *testing.T) {
		// Must not panic.
		animate(nil)
	})
}
