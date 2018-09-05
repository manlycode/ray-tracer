package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuple(t *testing.T) {
	tup := New(1, 2, 3)
	assert.Equal(t, Tuple{1, 2, 3}, tup)
}

func TestAddingTuples(t *testing.T) {
	a := New(3, -2, 5)
	b := New(-2, 3, 1)
	assert.Equal(t, New(1, 1, 6), Add(a, b))
}

func TestSubtractingTuples(t *testing.T) {
	a := New(3.5, 2.0, 1.0)
	b := New(5.5, 6.0, 7.0)

	assert.Equal(t, New(-2, -4, -6), Sub(a, b))
}

func TestNegatingTuples(t *testing.T) {
	a := New(1, -2, 3)
	assert.Equal(t, New(-1, 2, -3), Negate(a))
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
