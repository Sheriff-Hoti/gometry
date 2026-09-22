package main

import (
	"math"
	"testing"

	"github.com/gdamore/tcell/v3"
)

// textScreen records Put text per cell.
type textScreen struct {
	tcell.Screen
	cells map[[2]int]string
}

func (f *textScreen) Put(x, y int, text string, _ tcell.Style) (string, int) {
	if f.cells == nil {
		f.cells = make(map[[2]int]string)
	}
	f.cells[[2]int{x, y}] = text
	return "", 1
}

func TestDrawLabel(t *testing.T) {
	t.Run("writes text along the top row", func(t *testing.T) {
		f := &textScreen{}
		drawLabel(f, tcell.StyleDefault, "hi")
		if f.cells[[2]int{0, 0}] != "h" || f.cells[[2]int{1, 0}] != "i" {
			t.Errorf("got %v, want h at (0,0) and i at (1,0)", f.cells)
		}
		if len(f.cells) != 2 {
			t.Errorf("got %d cells, want 2", len(f.cells))
		}
	})

	t.Run("nil screen is a no-op", func(t *testing.T) {
		// Must not panic.
		drawLabel(nil, tcell.StyleDefault, "hi")
	})
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-12
}

func TestApplyZoom(t *testing.T) {
	t.Run("wheel up zooms in", func(t *testing.T) {
		if got := applyZoom(1.0, tcell.WheelUp); !almostEqual(got, 1.0*zoomStep) {
			t.Errorf("got %v, want %v", got, 1.0*zoomStep)
		}
	})

	t.Run("wheel down zooms out", func(t *testing.T) {
		if got := applyZoom(1.0, tcell.WheelDown); !almostEqual(got, 1.0/zoomStep) {
			t.Errorf("got %v, want %v", got, 1.0/zoomStep)
		}
	})

	t.Run("other buttons leave zoom unchanged", func(t *testing.T) {
		for _, b := range []tcell.ButtonMask{tcell.ButtonNone, tcell.Button1} {
			if got := applyZoom(2.0, b); got != 2.0 {
				t.Errorf("buttons %v: got %v, want 2.0", b, got)
			}
		}
	})

	t.Run("clamps at max zoom", func(t *testing.T) {
		if got := applyZoom(maxZoom, tcell.WheelUp); got != maxZoom {
			t.Errorf("got %v, want %v", got, maxZoom)
		}
	})

	t.Run("clamps at min zoom", func(t *testing.T) {
		if got := applyZoom(minZoom, tcell.WheelDown); got != minZoom {
			t.Errorf("got %v, want %v", got, minZoom)
		}
	})
}
