package main

import (
	"testing"

	"github.com/Sheriff-Hoti/gometry/shapes"
)

func TestDemosAllPlay(t *testing.T) {
	demos := newDemos()
	if len(demos) != 3 {
		t.Fatalf("len(demos) = %d, want 3 (one line per animation)", len(demos))
	}
	for i := range demos {
		if demos[i].step == nil {
			t.Fatalf("demos[%d].step is nil", i)
		}
	}

	before := make([]shapes.Line, len(demos))
	for i := range demos {
		before[i] = demos[i].line
	}
	for i := range demos {
		demos[i].step(&demos[i].line, 80, 24)
	}

	// Bounce line: B moves right, Y pinned.
	if demos[0].line.PointB.X != before[0].PointB.X+1 {
		t.Errorf("bounce B.X = %d, want %d", demos[0].line.PointB.X, before[0].PointB.X+1)
	}
	if demos[0].line.PointB.Y != before[0].PointB.Y {
		t.Errorf("bounce B.Y = %d, want unchanged %d", demos[0].line.PointB.Y, before[0].PointB.Y)
	}
	// Pendulum line: B moves down, X pinned.
	if demos[1].line.PointB.Y != before[1].PointB.Y+1 {
		t.Errorf("pendulum B.Y = %d, want %d", demos[1].line.PointB.Y, before[1].PointB.Y+1)
	}
	if demos[1].line.PointB.X != before[1].PointB.X {
		t.Errorf("pendulum B.X = %d, want unchanged %d", demos[1].line.PointB.X, before[1].PointB.X)
	}
	// Orbit line: pivot (58,12), radius 8, first 18° step lands on (66,14).
	if demos[2].line.PointB != (shapes.Point{X: 66, Y: 14}) {
		t.Errorf("orbit B = %v, want {66 14}", demos[2].line.PointB)
	}
	// No animation touches another line's anchor.
	for i := range demos {
		if demos[i].line.PointA != before[i].PointA {
			t.Errorf("demos[%d] PointA = %v, want unchanged %v", i, demos[i].line.PointA, before[i].PointA)
		}
	}
}
