// Demo showing hwo to compose types using struct embedding
package composer

import (
	"image/color"
	"math"
)

type Point struct { X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Distance as a Point method
func (p Point) Distance (q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
