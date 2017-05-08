package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"

	"github.com/jpoz/gascii"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "chris.png"
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	contentType := http.DetectContentType(f)
	buff := bytes.NewBuffer(f)
	var img image.Image

	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(buff)
		if err != nil {
			return
		}
	case "image/png":
		img, err = png.Decode(buff)
		if err != nil {
			return
		}
	default:
		return
	}

	config := gascii.Config{
		CanvasWidth: 80,
		HeightRatio: 1.8,
		CharPallette: []string{
			"M", "L", "i", "a", ",", "",
		},
	}

	g := gascii.Gascii{
		Config: &config,
		Image:  img,
	}

	g.Print()

	fmt.Println(g)
}
