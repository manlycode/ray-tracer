package tuple

import (
	"math"
)

// Tuple represents either a vector or a point
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func new(x float64, y float64, z float64, w float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: w}
}

//IsPoint returns true if the tuple is apoint
func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

//IsVector returns true when the tuple is a vector
func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

//Point returns a new point tuple
func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 1.0}
}

//Vector returns a new vector tuple
func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 0.0}
}

//Add adds two tuples and returns their resulting tuple
func Add(a Tuple, b Tuple) Tuple {
	return Tuple{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
		w: a.w + b.w,
	}
}

//Sub adds two tuples and returns their resulting tuple
func Sub(a Tuple, b Tuple) Tuple {
	return Tuple{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
		w: a.w - b.w,
	}
}

//Negate negates a tuple
func Negate(a Tuple) Tuple {
	return Sub(Vector(0, 0, 0), a)
}

//Mul multiplies a tuple by a factor
func Mul(a Tuple, scalar float64) Tuple {
	return Tuple{
		(scalar * a.x),
		(scalar * a.y),
		(scalar * a.z),
		(scalar * a.w),
	}
}

//Div divides a tuple by a factor
func Div(a Tuple, scalar float64) Tuple {
	return Tuple{
		(scalar / a.x),
		(scalar / a.y),
		(scalar / a.z),
		(scalar / a.w),
	}
}

//Magnitude calculates the magnitude of a Vector
func Magnitude(vec Tuple) float64 {
	return math.Sqrt(math.Pow(vec.x, 2) + math.Pow(vec.y, 2) + math.Pow(vec.z, 2))
}

//Normalize returns a normalized vector
func Normalize(vec Tuple) Tuple {
	mag := Magnitude(vec)
	return Vector(
		vec.x/mag,
		vec.y/mag,
		vec.z/mag,
	)
}
