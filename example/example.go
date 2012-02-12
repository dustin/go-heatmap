package main

import (
	"image"
	"image/png"
	"math/rand"
	"os"

	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
)

func main() {
	points := []heatmap.DataPoint{}
	for n := 0; n < 350; n++ {
		points = append(points,
			heatmap.P(rand.Float64(), rand.Float64()))
	}

	// scheme, _ := schemes.FromImage("../schemes/fire.png")
	scheme := schemes.Fire

	img := heatmap.Heatmap(image.Rect(0, 0, 1024, 1024),
		points, 150, 128, scheme)
	png.Encode(os.Stdout, img)
}
