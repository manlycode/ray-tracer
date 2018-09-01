package tuple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCast(t *testing.T) {
	pointTup := Tuple{1, 2, 3, 1}
	assert.Equal(t, Point{Tuple{1, 2, 3, 1}}, Cast(pointTup))

	vecTup := Tuple{1, 2, 3, 0}
	assert.Equal(t, Vector{Tuple{1, 2, 3, 0}}, Cast(vecTup))
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2, 3)
	assert.Equal(t, Tuple{1, 2, 3, 1}, p.GetTuple())
}

func TestNewVector(t *testing.T) {
	p := NewVector(1, 2, 3)
	assert.Equal(t, Tuple{1, 2, 3, 0}, p.GetTuple())
}

func TestAddingTuples(t *testing.T) {
	a := NewPoint(3, -2, 5)
	b := NewVector(-2, 3, 1)
	assert.Equal(t, NewPoint(1, 1, 6), Add(a, b))
}

func TestSubtractingPointFromPoint(t *testing.T) {
	a := NewPoint(3, 2, 1)
	b := NewPoint(5, 6, 7)

	assert.Equal(t, NewVector(-2, -4, -6), Sub(a, b))
}

func TestSubtractingPointFromNewVector(t *testing.T) {
	a := NewPoint(3, 2, 1)
	b := NewVector(5, 6, 7)
	assert.Equal(t, NewPoint(-2, -4, -6), Sub(a, b))
}

func TestSubtractingNewVectorFromNewVector(t *testing.T) {
	a := NewVector(3, 2, 1)
	b := NewVector(5, 6, 7)
	assert.Equal(t, NewVector(-2, -4, -6), Sub(a, b))
}

func TestNegatingTuples(t *testing.T) {
	a := NewVector(1, -2, 3)
	assert.Equal(t, NewVector(-1, 2, -3), Negate(a))
}

func TestMultiplyingTupleByScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.Equal(t, Tuple{3.5, -7, 10.5, -14.0}, Mul(a, 3.5))
}

func TestMultiplyingTupleByFraction(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2.0}, Mul(a, 0.5))
}

func TestDividingTupleByScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2.0}, Div(a, 2.0))
}

func TestMagnitudeOfNewVector(t *testing.T) {
	a := NewVector(1, 0, 0)
	assert.Equal(t, 1.0, a.Magnitude())

	a = NewVector(0, 1, 0)
	assert.Equal(t, 1.0, a.Magnitude())

	a = NewVector(0, 0, 1)
	assert.Equal(t, 1.0, a.Magnitude())

	a = NewVector(1, 2, 3)
	assert.Equal(t, math.Sqrt(14), a.Magnitude())

	a = NewVector(-1, -2, -3)
	assert.Equal(t, math.Sqrt(14), a.Magnitude())
}

func TestNormalize(t *testing.T) {
	a := NewVector(4, 0, 0)
	assert.Equal(t, NewVector(1, 0, 0), a.Normalize())

	a = NewVector(1, 2, 3)
	result := a.Normalize()
	assert.Equal(t, NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), result)
	assert.Equal(t, 1.0, result.Magnitude())
}

func TestDotProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	result := Dot(a, b)
	assert.Equal(t, 20.0, result)
}

func TestCrossProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	resultA := Cross(a, b)
	assert.Equal(t, NewVector(-1, 2, -1), resultA)

	resultB := Cross(b, a)
	assert.Equal(t, NewVector(1, -2, 1), resultB)
}
