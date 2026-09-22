package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Sheriff-Hoti/gometry/animations"
	"github.com/Sheriff-Hoti/gometry/shapes"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

// Zoom bounds and step for mouse-wheel zooming.
const (
	minZoom  = 0.2
	maxZoom  = 5.0
	zoomStep = 1.15
)

// applyZoom returns the zoom factor after one mouse event: wheel up zooms
// in, wheel down zooms out, anything else leaves it unchanged. The result
// stays within [minZoom, maxZoom].
func applyZoom(zoom float64, buttons tcell.ButtonMask) float64 {
	switch {
	case buttons&tcell.WheelUp != 0:
		zoom *= zoomStep
	case buttons&tcell.WheelDown != 0:
		zoom /= zoomStep
	}
	return min(max(zoom, minZoom), maxZoom)
}

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
	s.EnableMouse()
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

	zoom := 1.0

	// Event loop
	for {
		// Project the 12 edges to pixel space and draw them. Focal is
		// the pixel height, center is the pixel screen middle.
		w, h := s.Size()
		s.Clear()
		focal := float64(2*h) * zoom
		for _, l := range cube.Lines(focal, float64(w)/2, float64(h)) {
			l.Draw(s, cubeStyle)
		}
		drawLabel(s, defStyle, fmt.Sprintf("spinning cube (wheel zooms %.1fx, Esc quits)", zoom))
		s.Show()

		select {
		case <-ticker.C:
			spin(&cube)
		case ev := <-s.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventMouse:
				if ev.Buttons()&(tcell.WheelUp|tcell.WheelDown) != 0 {
					zoom = applyZoom(zoom, ev.Buttons())
				}
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
