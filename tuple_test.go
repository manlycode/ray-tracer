package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTuple(t *testing.T) {
	myTuple := new(4.3, -4.2, 3.1, 1.0)
	assert.Equal(t, 4.3, myTuple.x)
	assert.Equal(t, -4.2, myTuple.y)
	assert.Equal(t, 3.1, myTuple.z)
	assert.Equal(t, 1.0, myTuple.w)
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

	assert.Equal(t, 1.0, tupleC.x)
	assert.Equal(t, 1.0, tupleC.y)
	assert.Equal(t, 6.0, tupleC.z)
	assert.Equal(t, 1.0, tupleC.w)
}

func TestSubtractingPointFromPoint(t *testing.T) {
	tupleA := Point(3, 2, 1)
	tupleB := Point(5, 6, 7)
	tupleC := Sub(tupleA, tupleB)

	assert.Equal(t, -2.0, tupleC.x)
	assert.Equal(t, -4.0, tupleC.y)
	assert.Equal(t, -6.0, tupleC.z)
	assert.Equal(t, 0.0, tupleC.w)
	assert.True(t, tupleC.IsVector())
}

func TestSubtractingPointFromVector(t *testing.T) {
	tupleA := Point(3, 2, 1)
	tupleB := Vector(5, 6, 7)
	tupleC := Sub(tupleA, tupleB)

	assert.Equal(t, -2.0, tupleC.x)
	assert.Equal(t, -4.0, tupleC.y)
	assert.Equal(t, -6.0, tupleC.z)
	assert.Equal(t, 1.0, tupleC.w)
	assert.True(t, tupleC.IsPoint())
}

func TestSubtractingVectorFromVector(t *testing.T) {
	tupleA := Vector(3, 2, 1)
	tupleB := Vector(5, 6, 7)
	tupleC := Sub(tupleA, tupleB)

	assert.Equal(t, -2.0, tupleC.x)
	assert.Equal(t, -4.0, tupleC.y)
	assert.Equal(t, -6.0, tupleC.z)
	assert.True(t, tupleC.IsVector())
}

func TestNegatingTuples(t *testing.T) {
	tupleA := Vector(1, -2, 3)
	result := Negate(tupleA)

	assert.Equal(t, -1.0, result.x)
	assert.Equal(t, 2.0, result.y)
	assert.Equal(t, -3.0, result.z)
	assert.True(t, result.IsVector())
}
