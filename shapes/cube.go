package shapes

// Cube is an axis-aligned cube moved to Center, rotated by RotX then RotY,
// with Size as the half-edge length. Rotation is fixed at build time;
// animating it means rebuilding per frame with new angles.
type Cube struct {
	Center     Vec3
	Size       float64
	RotX, RotY float64
}

// cubeEdges pairs vertex indices into the 12 edges: bottom square,
// top square, then the 4 pillars.
var cubeEdges = [12][2]int{
	{0, 1}, {1, 3}, {3, 2}, {2, 0},
	{4, 5}, {5, 7}, {7, 6}, {6, 4},
	{0, 4}, {1, 5}, {2, 6}, {3, 7},
}

// Vertices returns the 8 rotated, translated corners. Index bit i is the
// sign of the axis (bit 0 = X, bit 1 = Y, bit 2 = Z).
func (c Cube) Vertices() [8]Vec3 {
	var vs [8]Vec3
	for i := range vs {
		v := Vec3{
			X: c.Size * (float64(i&1)*2 - 1),
			Y: c.Size * (float64((i>>1)&1)*2 - 1),
			Z: c.Size * (float64((i>>2)&1)*2 - 1),
		}
		v = RotateX(v, c.RotX)
		v = RotateY(v, c.RotY)
		vs[i] = Vec3{X: v.X + c.Center.X, Y: v.Y + c.Center.Y, Z: v.Z + c.Center.Z}
	}
	return vs
}

// Lines projects the 12 edges to pixel-space lines ready for DrawHalf.
// Focal, cx, cy are passed to Project; edges with an endpoint behind the
// camera are skipped.
func (c Cube) Lines(focal, cx, cy float64) []Line {
	vs := c.Vertices()
	lines := make([]Line, 0, len(cubeEdges))
	for _, e := range cubeEdges {
		a, okA := Project(vs[e[0]], focal, cx, cy)
		b, okB := Project(vs[e[1]], focal, cx, cy)
		if !okA || !okB {
			continue
		}
		lines = append(lines, Line{PointA: a, PointB: b})
	}
	return lines
}
