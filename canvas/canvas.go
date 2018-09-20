package canvas

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/manlycode/ray-tracer/color"
	"os"
)

// Canvas represents a rectangular grid of pixels
type Canvas struct {
	Width  int
	Height int
	Pixels [][]color.Color
}

// New returns a new canvas of width and height
func New(width int, height int) Canvas {
	return WithColor(width, height, color.Black())
}

// WithColor returns a canvas where all pixels are the same color
func WithColor(width int, height int, c color.Color) Canvas {
	pixels := make([][]color.Color, height)
	for row := range pixels {
		pixels[row] = make([]color.Color, width)
		for col := range pixels[row] {
			pixels[row][col] = c
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
	if x >= 0 && y >= 0 && x < c.Width && y < c.Height {
		c.Pixels[y][x] = col
	}
}

// WritePixelFlipped writes a color to a pixel
func (c Canvas) WritePixelFlipped(x int, y int, col color.Color) {
	c.WritePixel(x, c.Height-y, col)
}

// ToPPM returns a PPM representation of the canvas
func (c Canvas) ToPPM() (buffer bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height))

	for _, row := range c.Pixels {
		lineValues := []string{}
		var lineBuffer bytes.Buffer

		for _, pixel := range row {
			lineValues = append(lineValues, pixel.ToPPMStrings()...)
		}

		for i, value := range lineValues {

			lineBuffer.WriteString(value)
			if i+1 == len(lineValues) {
				lineBuffer.WriteString("\n")
			} else {
				if lineBuffer.Len()+1+len(value) > 70 {
					lineBuffer.WriteString("\n")
					buffer.WriteString(lineBuffer.String())
					lineBuffer.Reset()
				} else {
					lineBuffer.WriteString(" ")
				}
			}
		}

		buffer.WriteString(lineBuffer.String())
	}

	return
}

// Save writes the canvas PPM out to a file
func (c Canvas) Save(path string) {
	f, _ := os.Create(path)
	w := bufio.NewWriter(f)
	buf := c.ToPPM()
	buf.WriteTo(w)
	w.Flush()
}
