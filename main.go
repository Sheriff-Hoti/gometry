package main

import (
	"log"
	"time"

	"github.com/Sheriff-Hoti/gometry/animations"
	"github.com/Sheriff-Hoti/gometry/shapes"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

const (
	lineAx, lineAy = 10, 16
	lineBx, lineBy = 30, 40
)

// tickInterval is the delay between animation steps.
const tickInterval = 100 * time.Millisecond

// demo bundles a line with the animation step driving it and the style
// drawing it.
type demo struct {
	line  shapes.Line
	step  animations.Step
	style tcell.Style
}

// drawLabel writes text in the top-left corner of the screen.
func drawLabel(s tcell.Screen, style tcell.Style, text string) {
	if s == nil {
		return
	}
	for i, r := range text {
		s.Put(i, 0, string(r), style)
	}
}

// newDemos builds one line per animation so bounce, pendulum, and orbit
// all play at once.
func newDemos() []demo {
	demos := []demo{
		{
			line: shapes.Line{
				PointA: shapes.Point{X: lineAx, Y: lineAy},
				PointB: shapes.Point{X: lineBx, Y: lineBy},
			},
			style: tcell.StyleDefault.Foreground(color.Aqua).Background(color.Reset),
		},
		{
			line: shapes.Line{
				PointA: shapes.Point{X: 4, Y: 3},
				PointB: shapes.Point{X: 24, Y: 8},
			},
			style: tcell.StyleDefault.Foreground(color.Yellow).Background(color.Reset),
		},
		{
			line: shapes.Line{
				PointA: shapes.Point{X: 58, Y: 12},
				PointB: shapes.Point{X: 66, Y: 12},
			},
			style: tcell.StyleDefault.Foreground(color.Green).Background(color.Reset),
		},
	}
	steps := []animations.Step{
		animations.Bounce(1),
		animations.Pendulum(6, 1),
		animations.Orbit(20),
	}
	for i := range demos {
		demos[i].step = steps[i]
	}
	return demos
}

func main() {
	defStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)

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

	demos := newDemos()

	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	// Event loop
	for {
		// Draw before Show so each frame is visible immediately.
		s.Clear()
		for i := range demos {
			demos[i].line.DrawHalf(s, demos[i].style)
		}
		drawLabel(s, defStyle, "bounce + pendulum + orbit (Esc to quit)")
		s.Show()

		select {
		case <-ticker.C:
			// One animation step per tick for every line. Steps run in
			// double-height pixel space (see DrawHalf): x pixels map 1:1
			// to cells, y pixels are doubled, hence the 2*h heights.
			w, h := s.Size()
			for i := range demos {
				demos[i].step(&demos[i].line, w, 2*h)
			}
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
