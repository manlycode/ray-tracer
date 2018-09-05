package vector

import (
	"math"

	"github.com/manlycode/ray-tracer/tuple"
)

// Vector represents a
type Vector struct {
	tuple.Tuple
}

// GetTuple returns the tuple for tuple.Provider interface
func (v Vector) GetTuple() tuple.Tuple {
	return v.Tuple
}

func (v Vector) x() float64 { return v.Tuple[0] }
func (v Vector) y() float64 { return v.Tuple[1] }
func (v Vector) z() float64 { return v.Tuple[2] }
func (v Vector) w() float64 { return v.Tuple[3] }

// New creates a new vector
func New(x float64, y float64, z float64) Vector {
	return Vector{tuple.New(x, y, z, 0.0)}
}

//Magnitude calculates the magnitude of a Vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x(), 2) + math.Pow(v.y(), 2) + math.Pow(v.z(), 2))
}

//Normalize returns a normalized vector
func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	return New(
		v.x()/mag,
		v.y()/mag,
		v.z()/mag,
	)
}

//Dot returns the dot product of 2 vectors
func Dot(a Vector, b Vector) float64 {
	return a.x()*b.x() +
		a.y()*b.y() +
		a.z()*b.z() +
		a.w()*b.w()
}

//Cross returns the cross product of 2 vectors
func Cross(a Vector, b Vector) Vector {
	return New(
		a.y()*b.z()-a.z()*b.y(),
		a.z()*b.x()-a.x()*b.z(),
		a.x()*b.y()-a.y()*b.x(),
	)
}
