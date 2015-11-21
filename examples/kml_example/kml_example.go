package main

import (
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"

	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
)

func main() {
	// Cluster a few points randomly around Lawrence station.
	lawrence := heatmap.P(-121.996158, 37.370713)

	points := []heatmap.DataPoint{lawrence}
	for n := 0; n < 35; n++ {
		points = append(points,
			heatmap.P(lawrence.X()+(rand.Float64()/100.0)-0.005,
				lawrence.Y()+(rand.Float64()/100.0)-0.005))
	}

	kmlout, err := os.Create("test.kml")
	if err != nil {
		log.Fatalf("Error creating kml file:  %v", err)
	}
	defer kmlout.Close()

	imgout, err := os.Create("test.png")
	if err != nil {
		log.Fatalf("Error creating image file:  %v", err)
	}
	defer imgout.Close()

	img, err := heatmap.KML(image.Rect(0, 0, 1024, 1024),
		points, 200, 128, schemes.AlphaFire, "test.png", kmlout)
	png.Encode(imgout, img)
}
