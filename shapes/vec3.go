package shapes

import "math"

// Vec3 is a point or vector in 3D camera space: X right, Y down-screen,
// Z away from the camera. The camera sits at the origin, so visible
// points have Z > 0.
type Vec3 struct {
	X, Y, Z float64
}

// RotateX rotates v around the X axis by a radians (right-hand rule).
func RotateX(v Vec3, a float64) Vec3 {
	c, s := math.Cos(a), math.Sin(a)
	return Vec3{X: v.X, Y: v.Y*c - v.Z*s, Z: v.Y*s + v.Z*c}
}

// RotateY rotates v around the Y axis by a radians (right-hand rule).
func RotateY(v Vec3, a float64) Vec3 {
	c, s := math.Cos(a), math.Sin(a)
	return Vec3{X: v.X*c + v.Z*s, Y: v.Y, Z: -v.X*s + v.Z*c}
}

// RotateZ rotates v around the Z axis by a radians (right-hand rule).
func RotateZ(v Vec3, a float64) Vec3 {
	c, s := math.Cos(a), math.Sin(a)
	return Vec3{X: v.X*c - v.Y*s, Y: v.X*s + v.Y*c, Z: v.Z}
}

// Project maps a camera-space point to integer pixel coords with a
// perspective divide: farther points (larger Z) land closer to the screen
// center (cx, cy). Focal is the focal length in pixels; larger values look
// flatter. It returns false for points behind the camera (Z <= 0), which
// cannot be projected.
func Project(v Vec3, focal, cx, cy float64) (Point, bool) {
	if v.Z <= 0 {
		return Point{}, false
	}
	return Point{
		X: int(math.Round(cx + v.X*focal/v.Z)),
		Y: int(math.Round(cy + v.Y*focal/v.Z)),
	}, true
}
