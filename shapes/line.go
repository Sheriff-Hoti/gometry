package shapes

import (
	"github.com/gdamore/tcell/v3"
)

type Line struct {
	PointA Point
	PointB Point
	Screen tcell.Screen
	Style  tcell.Style
}

type Drawable interface {
	Draw()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

// Points returns every cell along the line from PointA to PointB,
// using Bresenham's algorithm (whole numbers only, all directions).
func (l Line) Points() []Point {
	dx := abs(l.PointB.X - l.PointA.X)
	dy := -abs(l.PointB.Y - l.PointA.Y)
	sx := sign(l.PointB.X - l.PointA.X)
	sy := sign(l.PointB.Y - l.PointA.Y)
	err := dx + dy

	var pts []Point
	x, y := l.PointA.X, l.PointA.Y
	for {
		pts = append(pts, Point{X: x, Y: y})
		if x == l.PointB.X && y == l.PointB.Y {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x += sx
		}
		if e2 <= dx {
			err += dx
			y += sy
		}
	}
	return pts
}

func (l Line) Draw() {
	for _, p := range l.Points() {
		l.Screen.Put(p.X, p.Y, "█", l.Style)
	}
}
