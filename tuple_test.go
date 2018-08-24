package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTuple(t *testing.T) {
	myTuple := new(4.3, -4.2, 3.1, 1.0)
	assert.Equal(t, myTuple.x, 4.3)
	assert.Equal(t, myTuple.y, -4.2)
	assert.Equal(t, myTuple.z, 3.1)
	assert.Equal(t, myTuple.w, 1.0)
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

	assert.Equal(t, point.x, 4.3)
	assert.Equal(t, point.y, -4.2)
	assert.Equal(t, point.z, 3.1)
	assert.True(t, point.IsPoint())
}

func TestCreateVector(t *testing.T) {
	vector := Vector(4.3, -4.2, 3.1)

	assert.Equal(t, vector.x, 4.3)
	assert.Equal(t, vector.y, -4.2)
	assert.Equal(t, vector.z, 3.1)
	assert.True(t, vector.IsVector())
}
