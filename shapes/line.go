package shapes

import (
	"github.com/gdamore/tcell/v3"
)

type Line struct {
	PointA Point
	PointB Point
}

type Drawable interface {
	Draw(s tcell.Screen, style tcell.Style)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	default:
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Points returns every cell along the line from PointA to PointB,
// using Bresenham's algorithm (whole numbers only, all directions).
func (l Line) Points() []Point {
	dx := abs(l.PointB.X - l.PointA.X)
	dy := -abs(l.PointB.Y - l.PointA.Y)
	sx := sign(l.PointB.X - l.PointA.X)
	sy := sign(l.PointB.Y - l.PointA.Y)
	err := dx + dy

	pts := make([]Point, 0, max(dx, -dy)+1)
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

func (l Line) Draw(s tcell.Screen, style tcell.Style) {
	if s == nil {
		return
	}
	w, h := s.Size()
	for _, p := range l.Points() {
		if p.X < 0 || p.Y < 0 || p.X >= w || p.Y >= h {
			continue
		}
		s.Put(p.X, p.Y, "█", style)
	}
}

// DrawHalf draws the line treating coordinates as double-height pixels to
// compensate for terminal cells being ~2x taller than wide: point (x, y)
// lands on the upper half of cell (x, y/2) when y is even and on the lower
// half when y is odd. A cell with both halves lit is drawn as a full block
// (█). Points mapping outside the screen are dropped. Merging happens per
// call: pixels from separate DrawHalf calls sharing a cell overwrite,
// last call wins.
func (l Line) DrawHalf(s tcell.Screen, style tcell.Style) {
	if s == nil {
		return
	}
	w, h := s.Size()
	type halves struct{ top, bot bool }
	cells := make(map[Point]halves)
	for _, p := range l.Points() {
		if p.X < 0 || p.Y < 0 {
			continue
		}
		cx, cy := p.X, p.Y/2
		if cx >= w || cy >= h {
			continue
		}
		hb := cells[Point{X: cx, Y: cy}]
		if p.Y%2 == 0 {
			hb.top = true
		} else {
			hb.bot = true
		}
		cells[Point{X: cx, Y: cy}] = hb
	}
	for cell, hb := range cells {
		switch {
		case hb.top && hb.bot:
			s.Put(cell.X, cell.Y, "█", style)
		case hb.top:
			s.Put(cell.X, cell.Y, "▀", style)
		default:
			s.Put(cell.X, cell.Y, "▄", style)
		}
	}
}
