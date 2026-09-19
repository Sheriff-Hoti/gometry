package main

import (
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
