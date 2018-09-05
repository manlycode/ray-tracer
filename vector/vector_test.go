package vector_test

import (
	"math"
	"testing"

	"github.com/manlycode/ray-tracer/tuple"
	"github.com/manlycode/ray-tracer/vector"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := vector.New(1.0, 2.0, 3.0)
	assert.Equal(
		t,
		vector.Vector{
			tuple.New(1, 2, 3, 0),
		},
		v,
	)
}

func TestVectorAdd(t *testing.T) {
	a := vector.New(1, 2, 3)
	b := vector.New(1, 2, 3)

	assert.Equal(t, tuple.New(2, 4, 6, 0), tuple.Add(a, b))
}

func TestVectorMagnitude(t *testing.T) {
	a := vector.New(1, 0, 0)
	assert.Equal(t, 1.0, a.Magnitude())

	a = vector.New(0, 1, 0)
	assert.Equal(t, 1.0, a.Magnitude())

	a = vector.New(0, 0, 1)
	assert.Equal(t, 1.0, a.Magnitude())

	a = vector.New(1, 2, 3)
	assert.Equal(t, math.Sqrt(14), a.Magnitude())

	a = vector.New(-1, -2, -3)
	assert.Equal(t, math.Sqrt(14), a.Magnitude())
}

func TestNormalize(t *testing.T) {
	a := vector.New(4, 0, 0)
	assert.Equal(t, vector.New(1, 0, 0), a.Normalize())

	a = vector.New(1, 2, 3)
	result := a.Normalize()
	assert.Equal(t, vector.New(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), result)
	assert.Equal(t, 1.0, result.Magnitude())
}

func TestDotProduct(t *testing.T) {
	a := vector.New(1, 2, 3)
	b := vector.New(2, 3, 4)

	result := vector.Dot(a, b)
	assert.Equal(t, 20.0, result)
}

func TestCrossProduct(t *testing.T) {
	a := vector.New(1, 2, 3)
	b := vector.New(2, 3, 4)

	resultA := vector.Cross(a, b)
	assert.Equal(t, vector.New(-1, 2, -1), resultA)

	resultB := vector.Cross(b, a)
	assert.Equal(t, vector.New(1, -2, 1), resultB)
}
