package tuple

// Tuple represents either a vector or a point
type Tuple []float64

type Provider interface {
	GetTuple() Tuple
}

// GetTuple returns the tuple
func (t Tuple) GetTuple() Tuple {
	return t
}

// New builds a new tuple
func New(value ...float64) Tuple {
	return value
}

// Add Adds two tuple.Providers
func Add(a Provider, b Provider) Tuple {
	result := Tuple{}
	for i, val := range a.GetTuple() {
		result = append(result, val+b.GetTuple()[i])
	}
	return result
}

// Sub subtracts two tuple.Providers
func Sub(a Provider, b Provider) Tuple {
	result := Tuple{}
	for i, val := range a.GetTuple() {
		result = append(result, val-b.GetTuple()[i])
	}
	return result
}

func zero(a Provider) Tuple {
	result := Tuple{}
	for index := 0; index < len(a.GetTuple()); index++ {
		result = append(result, 0)
	}

	return result
}

//Negate negates a tuple
func Negate(a Provider) Tuple {
	return Sub(zero(a), a)
}

//Mul multiplies a tuple by a factor
func Mul(a Provider, scalar float64) Tuple {
	result := Tuple{}
	for _, val := range a.GetTuple() {
		result = append(result, val*scalar)
	}
	return result
}

//Div multiplies a tuple by a factor
func Div(a Provider, scalar float64) Tuple {
	result := Tuple{}
	for _, val := range a.GetTuple() {
		result = append(result, val/scalar)
	}
	return result
}
