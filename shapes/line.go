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
	Draw() any
}

func (l Line) Draw() {
	l.Screen.Put(l.PointA.X, l.PointA.Y, "k", l.Style)
}
