package gascii

// Config hold all the configuration needed to make a new Ascii image.
type Config struct {
	CanvasWidth  int64
	HeightRatio  float64
	CharPallette []string
}
