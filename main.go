package main

import (
	"log"
	"time"

	"github.com/Sheriff-Hoti/gometry/shapes"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

const (
	lineAx, lineAy = 10, 16
	lineBx, lineBy = 30, 40
)

// animator holds animation state across ticks. A plain animate(l *Line)
// function cannot bounce or oscillate on its own because it has no memory
// of direction or phase, so that state lives here instead.
type animator struct {
	dx int // horizontal direction of PointB: +1 right, -1 left
}

// step advances the animation one tick: PointB moves horizontally and
// bounces off the left and right screen edges. Width is the screen width
// in cells and is used for edge detection.
func (a *animator) step(l *shapes.Line, width int) {
	if a == nil || l == nil || width <= 1 {
		return
	}
	if a.dx == 0 {
		a.dx = 1
	}
	l.PointB.X += a.dx
	if l.PointB.X >= width-1 {
		l.PointB.X = width - 1
		a.dx = -1
	} else if l.PointB.X <= 0 {
		l.PointB.X = 0
		a.dx = 1
	}
}

func main() {
	defStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)
	lineStyle := tcell.StyleDefault.Foreground(color.Aqua).Background(color.Reset)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	demoLine := shapes.Line{
		PointA: shapes.Point{X: lineAx, Y: lineAy},
		PointB: shapes.Point{X: lineBx, Y: lineBy},
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	anim := &animator{dx: 1}

	// Event loop
	for {
		// Draw before Show so each frame is visible immediately.
		s.Clear()
		demoLine.Draw(s, lineStyle)
		s.Show()

		select {
		case <-ticker.C:
			// One animation step per second: PointB bounces sideways.
			w, _ := s.Size()
			anim.step(&demoLine, w)
		case ev := <-s.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return
				} else if ev.Key() == tcell.KeyCtrlL {
					s.Sync()
				}
			}
		}
	}
}
