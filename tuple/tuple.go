package tuple

import (
	"math"
)

// Tuple represents either a vector or a point
type Tuple struct {
	x, y, z, w float64
}

// GetTuple returns the tuple
func (t Tuple) GetTuple() Tuple {
	return t
}

// Provider is a struct that provides a tuple
type Provider interface {
	GetTuple() Tuple
}

// Cast returns a Point or a Vector depending on the w value
func Cast(p Provider) Provider {
	t := p.GetTuple()
	if t.w == 0.0 {
		return Vector{t}
	}

	if t.w == 1.0 {
		return Point{t}
	}

	return t
}

// Add Adds two tuple.Providers
func Add(aProv Provider, bProv Provider) Provider {
	a := aProv.GetTuple()
	b := bProv.GetTuple()
	tuple := Tuple{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
		w: a.w + b.w,
	}

	return Cast(tuple)
}

// Sub subtracts two tuple.Providers
func Sub(aProv Provider, bProv Provider) Provider {
	a := aProv.GetTuple()
	b := bProv.GetTuple()
	tuple := Tuple{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
		w: a.w - b.w,
	}

	return Cast(tuple)
}

//Negate negates a tuple
func Negate(a Provider) Provider {
	return Sub(NewVector(0, 0, 0), a)
}

//Mul multiplies a tuple by a factor
func Mul(aProv Provider, scalar float64) Provider {
	a := aProv.GetTuple()
	tuple := Tuple{
		(scalar * a.x),
		(scalar * a.y),
		(scalar * a.z),
		(scalar * a.w),
	}
	return Cast(tuple)
}

//Div divides a tuple by a factor
func Div(aProv Provider, scalar float64) Provider {
	a := aProv.GetTuple()
	tuple := Tuple{
		(a.x / scalar),
		(a.y / scalar),
		(a.z / scalar),
		(a.w / scalar),
	}

	return Cast(tuple)
}

// Point represents a point on a cartesian plane
type Point struct {
	Tuple
}

// NewPoint builds a new point
func NewPoint(x float64, y float64, z float64) Point {
	return Point{Tuple{x, y, z, 1.0}}
}

// Vector represents a point on a cartesian plane
type Vector struct {
	Tuple
}

// NewVector builds a new point
func NewVector(x float64, y float64, z float64) Vector {
	return Vector{Tuple{x, y, z, 0.0}}
}

//Magnitude calculates the magnitude of a Vector
func (vec Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.x, 2) + math.Pow(vec.y, 2) + math.Pow(vec.z, 2))
}

//Normalize returns a normalized vector
func (vec Vector) Normalize() Vector {
	tuple := vec.GetTuple()
	mag := vec.Magnitude()
	return NewVector(
		tuple.x/mag,
		tuple.y/mag,
		tuple.z/mag,
	)
}

//Dot returns the dot product of 2 vectors
func Dot(aVec Vector, bVec Vector) float64 {
	a := aVec.GetTuple()
	b := bVec.GetTuple()

	return a.x*b.x +
		a.y*b.y +
		a.z*b.z +
		a.w*b.w
}

//Cross returns the cross product of 2 vectors
func Cross(a Vector, b Vector) Vector {
	return NewVector(
		a.y*b.z-a.z*b.y,
		a.z*b.x-a.x*b.z,
		a.x*b.y-a.y*b.x,
	)
}
