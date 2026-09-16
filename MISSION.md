# Mission: Bresenham's line algorithm

## Why
The gometry terminal renderer draws lines with Bresenham (`shapes/line.go`), and I want to actually understand the code I ship — to debug it, review changes to it, and extend it with confidence.

## Success looks like
- I can hand-trace the error term for a shallow line and point at the exact lines of `Points()` doing each step
- I can explain how the same loop covers steep, reversed, and negative-direction lines
- I can predict what the edge cases (single point, perfectly vertical/horizontal) do without running the code

## Constraints
- Short lessons; one tangible win each
- Everything grounded in my own `shapes` package, not abstract examples

## Out of scope
- Bresenham circles/ellipses (a possible later mission, not this one)
- Floating-point line rasterization (e.g. Xiaolin Wu)
