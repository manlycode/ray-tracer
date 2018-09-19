package canvas_test

import (
	"testing"

	"github.com/manlycode/ray-tracer/canvas"
	"github.com/manlycode/ray-tracer/color"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := canvas.New(320, 240)
	assert.Equal(t, c.Width, 320)
	assert.Equal(t, c.Height, 240)

	for _, row := range c.Pixels {
		assert.Equal(t, 320, len(row))
		for i := range row {
			assert.Equal(t, color.Black(), row[i])
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := canvas.New(320, 240)
	c.WritePixel(10, 10, color.Red())
	assert.Equal(t, color.Red(), c.PixelAt(10, 10))
}

func TestToPPM(t *testing.T) {
	canv := canvas.New(5, 3)
	a := color.New(1.5, 0, 0)
	b := color.New(0, 0.5, 0)
	c := color.New(-0.5, 0, 1)

	canv.WritePixel(0, 0, a)
	canv.WritePixel(2, 1, b)
	canv.WritePixel(4, 2, c)

	ppmOutput := canv.ToPPM()

	expected := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	assert.Equal(t, expected, ppmOutput.String())
}

func TestToPPMLineLength(t *testing.T) {
	a := color.New(1, 0.8, 0.6)
	canv := canvas.WithColor(10, 2, a)

	ppmOutput := canv.ToPPM()

	expected := `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153 
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255
204 153 255 204 153 255 204 153 255 204 153 255 204 153 
`

	assert.Equal(t, expected, ppmOutput.String())
}
