package point_test

import (
	"testing"

	"github.com/manlycode/ray-tracer/point"
	"github.com/manlycode/ray-tracer/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := point.New(1.0, 2.0, 3.0)
	assert.Equal(
		t,
		point.Point{
			tuple.New(1, 2, 3, 1),
		},
		v,
	)
}
