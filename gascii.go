package gascii

import (
	"fmt"
	"image"
	"image/color"
)

// Gascii is the configuration and the image
type Gascii struct {
	Config *Config
	Image  image.Image
}

// Print out
func (g Gascii) Print() error {
	b := g.Image.Bounds()

	sliceSize := int(int64(b.Max.X) / g.Config.CanvasWidth)
	sliceHeight := int(float64(sliceSize) * g.Config.HeightRatio)
	fmt.Println(sliceSize)

	colorStep := 255 / (len(g.Config.CharPallette))
	fmt.Println(colorStep)

	for y := b.Min.Y; y < b.Max.Y; y += sliceHeight {
		for x := b.Min.X; x < b.Max.X; x += sliceSize {

			color := g.Image.At(x, y)
			grayColor := convertToGray(color)

			paletteIndex := grayColor.Y / uint8(colorStep)

			// fmt.Print(paletteIndex)
			fmt.Print(g.Config.CharPallette[paletteIndex])
		}
		fmt.Print("\n")
	}

	return nil
}

// From http://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
func convertToGray(c color.Color) color.Gray {
	r, g, b, _ := c.RGBA()
	lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	return color.Gray{uint8(lum / 256)}
}
