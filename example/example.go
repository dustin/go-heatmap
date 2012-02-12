package main

import (
	"image"
	"image/png"
	"log"
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

	scheme, err := heatmap.SchemeFromImage("../fire.png")
	if err != nil {
		log.Fatalf("Error initializing scheme: %v", err)
	}

	img := heatmap.Heatmap(image.Rect(0, 0, 1024, 1024),
		points, 150, 128, scheme)
	png.Encode(os.Stdout, img)
}
