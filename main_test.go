package main

import (
	"testing"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

func TestAnimatorStep(t *testing.T) {
	t.Run("moves point B one cell right by default", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 10, Y: 16},
			PointB: shapes.Point{X: 30, Y: 40},
		}
		a := &animator{}
		a.step(&l, 80)
		if l.PointB.X != 31 {
			t.Errorf("PointB.X = %d, want 31", l.PointB.X)
		}
		if l.PointB.Y != 40 {
			t.Errorf("PointB.Y = %d, want 40 (unchanged)", l.PointB.Y)
		}
		if l.PointA != (shapes.Point{X: 10, Y: 16}) {
			t.Errorf("PointA = %v, want unchanged {10 16}", l.PointA)
		}
		if a.dx != 1 {
			t.Errorf("dx = %d, want 1", a.dx)
		}
	})

	t.Run("bounces off the right edge", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 0, Y: 0},
			PointB: shapes.Point{X: 8, Y: 0},
		}
		a := &animator{dx: 1}
		a.step(&l, 10)
		if l.PointB.X != 9 {
			t.Fatalf("PointB.X = %d, want 9 (clamped to edge)", l.PointB.X)
		}
		if a.dx != -1 {
			t.Fatalf("dx = %d, want -1 after bounce", a.dx)
		}
		a.step(&l, 10)
		if l.PointB.X != 8 {
			t.Errorf("PointB.X = %d, want 8 (moving left)", l.PointB.X)
		}
	})

	t.Run("bounces off the left edge", func(t *testing.T) {
		l := shapes.Line{
			PointA: shapes.Point{X: 5, Y: 5},
			PointB: shapes.Point{X: 1, Y: 5},
		}
		a := &animator{dx: -1}
		a.step(&l, 10)
		if l.PointB.X != 0 {
			t.Fatalf("PointB.X = %d, want 0 (clamped to edge)", l.PointB.X)
		}
		if a.dx != 1 {
			t.Fatalf("dx = %d, want 1 after bounce", a.dx)
		}
	})

	t.Run("nil animator or line is a no-op", func(t *testing.T) {
		// Must not panic.
		var a *animator
		a.step(&shapes.Line{}, 80)
		(&animator{}).step(nil, 80)
		(&animator{dx: 1}).step(&shapes.Line{}, 0)
	})
}
