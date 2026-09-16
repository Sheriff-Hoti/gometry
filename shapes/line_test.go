package shapes

import (
	"reflect"
	"testing"

	"github.com/gdamore/tcell/v3"
)

// fakeScreen records Put calls while reporting a fixed size.
type fakeScreen struct {
	tcell.Screen
	w, h int
	puts []Point
}

func (f *fakeScreen) Size() (int, int) { return f.w, f.h }

func (f *fakeScreen) Put(x, y int, _ string, _ tcell.Style) (string, int) {
	f.puts = append(f.puts, Point{X: x, Y: y})
	return "", 1
}

func TestPoints(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want []Point
	}{
		{
			name: "shallow up",
			a:    Point{X: 0, Y: 0},
			b:    Point{X: 5, Y: 2},
			want: []Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 2}, {X: 5, Y: 2}},
		},
		{
			name: "shallow down",
			a:    Point{X: 0, Y: 2},
			b:    Point{X: 5, Y: 0},
			want: []Point{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 0}, {X: 5, Y: 0}},
		},
		{
			name: "steep up",
			a:    Point{X: 0, Y: 0},
			b:    Point{X: 2, Y: 5},
			want: []Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 2}, {X: 1, Y: 3}, {X: 2, Y: 4}, {X: 2, Y: 5}},
		},
		{
			name: "steep down",
			a:    Point{X: 0, Y: 5},
			b:    Point{X: 2, Y: 0},
			want: []Point{{X: 0, Y: 5}, {X: 0, Y: 4}, {X: 1, Y: 3}, {X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 0}},
		},
		{
			name: "horizontal",
			a:    Point{X: 1, Y: 3},
			b:    Point{X: 4, Y: 3},
			want: []Point{{X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 4, Y: 3}},
		},
		{
			name: "horizontal reversed",
			a:    Point{X: 4, Y: 3},
			b:    Point{X: 1, Y: 3},
			want: []Point{{X: 4, Y: 3}, {X: 3, Y: 3}, {X: 2, Y: 3}, {X: 1, Y: 3}},
		},
		{
			name: "vertical",
			a:    Point{X: 2, Y: 0},
			b:    Point{X: 2, Y: 3},
			want: []Point{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},
		},
		{
			name: "vertical reversed",
			a:    Point{X: 2, Y: 3},
			b:    Point{X: 2, Y: 0},
			want: []Point{{X: 2, Y: 3}, {X: 2, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 0}},
		},
		{
			name: "diagonal",
			a:    Point{X: 0, Y: 0},
			b:    Point{X: 3, Y: 3},
			want: []Point{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
		},
		{
			name: "negative coords",
			a:    Point{X: -2, Y: 0},
			b:    Point{X: 0, Y: 0},
			want: []Point{{X: -2, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 0}},
		},
		{
			name: "single point",
			a:    Point{X: 4, Y: 4},
			b:    Point{X: 4, Y: 4},
			want: []Point{{X: 4, Y: 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Line{PointA: tt.a, PointB: tt.b}.Points()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDraw(t *testing.T) {
	t.Run("nil screen is a no-op", func(t *testing.T) {
		// Must not panic.
		Line{PointA: Point{X: 0, Y: 0}, PointB: Point{X: 3, Y: 0}}.Draw(nil, tcell.StyleDefault)
	})

	t.Run("plots every point in bounds", func(t *testing.T) {
		f := &fakeScreen{w: 10, h: 10}
		Line{PointA: Point{X: 1, Y: 1}, PointB: Point{X: 3, Y: 1}}.Draw(f, tcell.StyleDefault)
		want := []Point{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}}
		if !reflect.DeepEqual(f.puts, want) {
			t.Errorf("got %v, want %v", f.puts, want)
		}
	})

	t.Run("clips out of bounds points", func(t *testing.T) {
		f := &fakeScreen{w: 5, h: 5}
		Line{PointA: Point{X: 0, Y: 0}, PointB: Point{X: 10, Y: 10}}.Draw(f, tcell.StyleDefault)
		want := []Point{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 4}}
		if !reflect.DeepEqual(f.puts, want) {
			t.Errorf("got %v, want %v", f.puts, want)
		}
	})
}
