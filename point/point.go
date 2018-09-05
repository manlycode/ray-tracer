package point

import (
	"github.com/manlycode/ray-tracer/tuple"
)

// Point represents a point in cartesian space
type Point struct {
	tuple.Tuple
}

// GetTuple returns the tuple for the point
// to satisfy the tuple.Provider interface
func (p Point) GetTuple() tuple.Tuple {
	return p.Tuple
}

// New returns a new Point
func New(x float64, y float64, z float64) Point {
	return Point{tuple.Tuple{x, y, z, 1.0}}
}
