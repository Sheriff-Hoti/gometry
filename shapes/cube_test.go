package shapes

import (
	"math"
	"testing"
)

func approxVec(got, want Vec3, t *testing.T) {
	t.Helper()
	d := math.Hypot(got.X-want.X, math.Hypot(got.Y-want.Y, got.Z-want.Z))
	if d > 1e-9 {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRotate(t *testing.T) {
	t.Run("quarter turn about Y sends +X to -Z", func(t *testing.T) {
		approxVec(RotateY(Vec3{X: 1}, math.Pi/2), Vec3{Z: -1}, t)
	})

	t.Run("quarter turn about X sends +Y to +Z", func(t *testing.T) {
		approxVec(RotateX(Vec3{Y: 1}, math.Pi/2), Vec3{Z: 1}, t)
	})

	t.Run("quarter turn about Z sends +X to +Y", func(t *testing.T) {
		approxVec(RotateZ(Vec3{X: 1}, math.Pi/2), Vec3{Y: 1}, t)
	})

	t.Run("full turn returns to start", func(t *testing.T) {
		v := Vec3{X: 3, Y: -2, Z: 5}
		approxVec(RotateY(RotateX(RotateZ(v, 2*math.Pi), 2*math.Pi), 2*math.Pi), v, t)
	})
}

func TestProject(t *testing.T) {
	t.Run("centered point lands on center", func(t *testing.T) {
		p, ok := Project(Vec3{Z: 10}, 10, 40, 24)
		if !ok || p != (Point{X: 40, Y: 24}) {
			t.Errorf("got %v, %v; want {40 24}, true", p, ok)
		}
	})

	t.Run("farther points shrink toward center", func(t *testing.T) {
		near, _ := Project(Vec3{X: 10, Z: 10}, 10, 40, 24)
		far, _ := Project(Vec3{X: 10, Z: 20}, 10, 40, 24)
		if near != (Point{X: 50, Y: 24}) || far != (Point{X: 45, Y: 24}) {
			t.Errorf("got near=%v far=%v, want {50 24} {45 24}", near, far)
		}
	})

	t.Run("points behind the camera are rejected", func(t *testing.T) {
		for _, z := range []float64{0, -1} {
			if _, ok := Project(Vec3{Z: z}, 10, 40, 24); ok {
				t.Errorf("Z=%v: got ok=true, want false", z)
			}
		}
	})
}

func TestCube(t *testing.T) {
	t.Run("unrotated vertices sit on the corners", func(t *testing.T) {
		vs := Cube{Center: Vec3{Z: 5}, Size: 1}.Vertices()
		approxVec(vs[0], Vec3{X: -1, Y: -1, Z: 4}, t)
		approxVec(vs[7], Vec3{X: 1, Y: 1, Z: 6}, t)
	})

	t.Run("front-facing cube projects 12 edges", func(t *testing.T) {
		lines := Cube{Center: Vec3{Z: 32}, Size: 10, RotX: 0.35, RotY: 0.5}.Lines(48, 40, 24)
		if len(lines) != 12 {
			t.Fatalf("got %d lines, want 12", len(lines))
		}
		for i, l := range lines {
			if l.PointA == l.PointB {
				t.Errorf("line %d is degenerate: %v", i, l)
			}
		}
	})

	t.Run("cube behind the camera projects nothing", func(t *testing.T) {
		lines := Cube{Center: Vec3{Z: -5}, Size: 1}.Lines(48, 40, 24)
		if len(lines) != 0 {
			t.Errorf("got %d lines, want 0", len(lines))
		}
	})
}
