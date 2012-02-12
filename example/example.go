package main

import (
	"image"
	"image/png"
	"math/rand"
	"os"

	"github.com/dustin/go-heatmap"
)

func main() {
	points := []heatmap.DataPoint{}
	for n := 0; n < 200; n++ {
		points = append(points,
			heatmap.DataPoint{rand.Float64(), rand.Float64()})
	}

	img := heatmap.Heatmap(image.Rect(0, 0, 1024, 1024),
		points, 150, 128, heatmap.FireScheme)
	png.Encode(os.Stdout, img)
}
