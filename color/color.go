package color

import (
	"github.com/manlycode/ray-tracer/tuple"
)

// Color represents a color
type Color struct {
	tuple.Tuple
}

// GetTuple makes Color conform to the tuple.Provider interface
func (c Color) GetTuple() tuple.Tuple {
	return c.Tuple
}

func (c Color) r() float64 { return c.Tuple[0] }
func (c Color) g() float64 { return c.Tuple[1] }
func (c Color) b() float64 { return c.Tuple[2] }

// New creates a new color
func New(r float64, g float64, b float64) Color {
	return Color{tuple.Tuple{r, g, b}}
}

// Black returns the representation of Black
func Black() Color {
	return New(0, 0, 0)
}

// Red returns the representation of Black
func Red() Color {
	return New(1.0, 0, 0)
}

// Add adds colors
func Add(a Color, b Color) Color {
	return Color{tuple.Add(a, b)}
}

// Sub adds colors
func Sub(a Color, b Color) Color {
	return Color{tuple.Sub(a, b)}
}

// Mul adds colors
func Mul(a Color, scalar float64) Color {
	return Color{tuple.Mul(a, scalar)}
}

// HadamardProduct subtracts two tuple.Providers
func HadamardProduct(a Color, b Color) Color {
	result := tuple.Tuple{}
	for i, val := range a.GetTuple() {
		tup := b.GetTuple()
		bVal := tup[i]
		result = append(result, val*bVal)
	}
	return Color{result}
}
