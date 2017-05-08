package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/jpoz/gascii"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./chris.png")
	check(err)

	image, err := png.Decode(f)
	check(err)

	config := gascii.Config{
		CanvasWidth: 80,
		HeightRatio: 1.8,
		CharPallette: []string{
			"M", "L", "i", "a", ",", "",
		},
	}

	g := gascii.Gascii{
		Config: &config,
		Image:  image,
	}

	g.Print()

	fmt.Println(g)
}
