package tuple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertTuplesEqual(t *testing.T, a Tuple, b Tuple) {
	assert.Equal(t, a.x, b.x, "x values are not equal")
	assert.Equal(t, a.y, b.y, "y values are not equal")
	assert.Equal(t, a.z, b.z, "z values are not equal")
	assert.Equal(t, a.w, b.w, "w values are not equal")
}

func TestCreateTuple(t *testing.T) {
	myTuple := new(4.3, -4.2, 3.1, 1.0)
	assertTuplesEqual(t, Point(4.3, -4.2, 3.1), myTuple)
}

func TestIsPoint(t *testing.T) {
	pointTuple := new(4.3, -4.2, 3.1, 1.0)
	assert.True(t, pointTuple.IsPoint())

	vectorTuple := new(4.3, -4.2, 3.1, 0.0)
	assert.False(t, vectorTuple.IsPoint())
}

func TestIsVector(t *testing.T) {
	pointTuple := new(4.3, -4.2, 3.1, 1.0)
	assert.False(t, pointTuple.IsVector())

	vectorTuple := new(4.3, -4.2, 3.1, 0.0)
	assert.True(t, vectorTuple.IsVector())
}

func TestCreatePoint(t *testing.T) {
	point := Point(4.3, -4.2, 3.1)

	assert.Equal(t, 4.3, point.x)
	assert.Equal(t, -4.2, point.y)
	assert.Equal(t, 3.1, point.z)
	assert.True(t, point.IsPoint())
}

func TestCreateVector(t *testing.T) {
	vector := Vector(4.3, -4.2, 3.1)

	assert.Equal(t, 4.3, vector.x)
	assert.Equal(t, -4.2, vector.y)
	assert.Equal(t, 3.1, vector.z)
	assert.True(t, vector.IsVector())
}

func TestAddingTuples(t *testing.T) {
	tupleA := new(3, -2, 5, 1)
	tupleB := new(-2, 3, 1, 0)
	tupleC := Add(tupleA, tupleB)

	assertTuplesEqual(t, Point(1.0, 1.0, 6.0), tupleC)
}

func TestSubtractingPointFromPoint(t *testing.T) {
	tupleA := Point(3, 2, 1)
	tupleB := Point(5, 6, 7)
	tupleC := Sub(tupleA, tupleB)

	assertTuplesEqual(t, Vector(-2, -4, -6), tupleC)
}

func TestSubtractingPointFromVector(t *testing.T) {
	tupleA := Point(3, 2, 1)
	tupleB := Vector(5, 6, 7)
	tupleC := Sub(tupleA, tupleB)
	assertTuplesEqual(t, Point(-2, -4, -6), tupleC)
}

func TestSubtractingVectorFromVector(t *testing.T) {
	tupleA := Vector(3, 2, 1)
	tupleB := Vector(5, 6, 7)
	tupleC := Sub(tupleA, tupleB)
	assertTuplesEqual(t, Vector(-2, -4, -6), tupleC)
}

func TestNegatingTuples(t *testing.T) {
	tupleA := Vector(1, -2, 3)
	result := Negate(tupleA)
	assertTuplesEqual(t, Vector(-1, 2, -3), result)
}

func TestMultiplyingTupleByScalar(t *testing.T) {
	tupleA := new(1, -2, 3, -4)
	result := Mul(tupleA, 3.5)

	assertTuplesEqual(t, new(3.5, -7, 10.5, -14.0), result)
}

func TestMultiplyingTupleByFraction(t *testing.T) {
	tupleA := new(1, -2, 3, -4)
	result := Mul(tupleA, 0.5)
	assertTuplesEqual(t, new(0.5, -1, 1.5, -2.0), result)
}

func TestDividingTupleByScalar(t *testing.T) {
	tupleA := new(1, -2, 3, -4)
	result := Div(tupleA, 2.0)

	assertTuplesEqual(t, new(0.5, -1, 1.5, -2.0), result)
}

func TestMagnitudeOfVector(t *testing.T) {
	tupleA := Vector(1, 0, 0)
	result := Magnitude(tupleA)
	assert.Equal(t, 1.0, result)

	tupleA = Vector(0, 1, 0)
	result = Magnitude(tupleA)
	assert.Equal(t, 1.0, result)

	tupleA = Vector(0, 0, 1)
	result = Magnitude(tupleA)
	assert.Equal(t, 1.0, result)

	tupleA = Vector(1, 2, 3)
	result = Magnitude(tupleA)
	assert.Equal(t, math.Sqrt(14), result)

	tupleA = Vector(-1, -2, -3)
	result = Magnitude(tupleA)
	assert.Equal(t, math.Sqrt(14), result)
}

func TestNormalize(t *testing.T) {
	tupleA := Vector(4, 0, 0)
	result := Normalize(tupleA)
	assertTuplesEqual(t, Vector(1, 0, 0), result)

	tupleA = Vector(1, 2, 3)
	result = Normalize(tupleA)
	assertTuplesEqual(t, Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), result)
	assert.Equal(t, 1.0, Magnitude(result))
}

func TestDotProduct(t *testing.T) {
	tupleA := Vector(1, 2, 3)
	tupleB := Vector(2, 3, 4)
	result := Dot(tupleA, tupleB)
	assert.Equal(t, 20.0, result)
}

func TestCrossProduct(t *testing.T) {
	tupleA := Vector(1, 2, 3)
	tupleB := Vector(2, 3, 4)

	resultA := Cross(tupleA, tupleB)
	assert.Equal(t, Vector(-1, 2, -1), resultA)

	resultB := Cross(tupleB, tupleA)
	assert.Equal(t, Vector(1, -2, 1), resultB)
}
