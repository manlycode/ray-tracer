package color_test

import (
	"testing"

	"github.com/manlycode/ray-tracer/color"
	"github.com/manlycode/ray-tracer/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := color.New(0, 0, 0)
	assert.Equal(
		t,
		color.Color{
			tuple.New(0, 0, 0),
		},
		c,
	)
}

func TestAdd(t *testing.T) {
	a := color.New(0.9, 0.6, 0.75)
	b := color.New(0.7, 0.1, 0.25)

	assert.Equal(t, color.New(1.6, 0.7, 1.0), color.Add(a, b))
}

func TestSub(t *testing.T) {
	a := color.New(0.9, 0.6, 0.75)
	b := color.New(0.7, 0.1, 0.25)

	//WTF?
	assert.Equal(t, color.New(0.2300000000000007, 0.5, 0.5), color.Sub(a, b))
}

func TestMul(t *testing.T) {
	a := color.New(0.2, 0.3, 0.4)
	assert.Equal(t, color.New(0.4, 0.6, 0.8), color.Mul(a, 2))
}

func TestBlend(t *testing.T) {
	a := color.New(1, 0.2, 0.4)
	b := color.New(0.9, 1, 0.1)

	assert.Equal(t, color.New(0.9, 0.2, 0.04000000000000001), color.HadamardProduct(a, b))
}
