package canvas

import (
	"fmt"
	"github.com/manlycode/ray-tracer/color"
)

// Canvas represents a rectangular grid of pixels
type Canvas struct {
	Width  int
	Height int
	Pixels [][]color.Color
}

// New returns a new canvas of width and height
func New(width int, height int) Canvas {
	pixels := make([][]color.Color, height)
	for row := range pixels {
		pixels[row] = make([]color.Color, width)
		for col := range pixels[row] {
			pixels[row][col] = color.Black()
		}
	}

	return Canvas{width, height, pixels}
}

// PixelAt returns the color for the pixel at x and y
func (c Canvas) PixelAt(x int, y int) color.Color {
	return c.Pixels[y][x]
}

// WritePixel writes a color to a pixel
func (c Canvas) WritePixel(x int, y int, col color.Color) {
	c.Pixels[y][x] = col
}

// ToPPM
func (c Canvas) ToPPM() (ppm string) {
	ppm = fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
	return
}
