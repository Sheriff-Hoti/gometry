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

// animate advances the line one animation step by moving PointB one cell
// to the right. It takes a pointer so the caller's line is updated.
func animate(l *shapes.Line) {
	if l == nil {
		return
	}
	l.PointB.X++
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

	// Event loop
	for {
		// Draw before Show so each frame is visible immediately.
		s.Clear()
		demoLine.Draw(s, lineStyle)
		s.Show()

		select {
		case <-ticker.C:
			// One animation step per second: PointB moves right.
			animate(&demoLine)
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
