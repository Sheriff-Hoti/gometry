package shapes

import (
	"reflect"
	"testing"
)

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
			name: "horizontal",
			a:    Point{X: 1, Y: 3},
			b:    Point{X: 4, Y: 3},
			want: []Point{{X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 4, Y: 3}},
		},
		{
			name: "vertical",
			a:    Point{X: 2, Y: 0},
			b:    Point{X: 2, Y: 3},
			want: []Point{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},
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
