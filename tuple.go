package tuple

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
