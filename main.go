package main

import (
	"log"
	"time"

	"github.com/Sheriff-Hoti/gometry/animations"
	"github.com/Sheriff-Hoti/gometry/shapes"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

// tickInterval is the delay between animation steps.
const tickInterval = 60 * time.Millisecond

// drawLabel writes text in the top-left corner of the screen.
func drawLabel(s tcell.Screen, style tcell.Style, text string) {
	if s == nil {
		return
	}
	for i, r := range text {
		s.Put(i, 0, string(r), style)
	}
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

	// A static cube tilted so all three faces show. Depth 32 with the
	// focal length tracking screen height keeps it ~30px wide on an
	// 80x24 terminal; Draw clips whatever falls outside.
	cube := shapes.Cube{
		Center: shapes.Vec3{Z: 32},
		Size:   10,
		RotX:   0.35,
		RotY:   0.5,
	}
	cubeStyle := tcell.StyleDefault.Foreground(color.Aqua).Background(color.Reset)
	spin := animations.Spin(0.06, 0.1)

	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	// Event loop
	for {
		// Project the 12 edges to pixel space and draw them. Focal is
		// the pixel height, center is the pixel screen middle.
		w, h := s.Size()
		s.Clear()
		focal := float64(2 * h)
		for _, l := range cube.Lines(focal, float64(w)/2, float64(h)) {
			l.Draw(s, cubeStyle)
		}
		drawLabel(s, defStyle, "spinning cube (Esc to quit)")
		s.Show()

		select {
		case <-ticker.C:
			spin(&cube)
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
