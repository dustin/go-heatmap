package main

import (
	"image"
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

	kmzout, err := os.Create("test.kmz")
	if err != nil {
		log.Fatalf("Error creating kml file:  %v", err)
	}
	defer kmzout.Close()
	err = heatmap.KMZ(image.Rect(0, 0, 1024, 1024),
		points, 200, 128, schemes.AlphaFire, kmzout)
	if err != nil {
		log.Fatalf("Error creating heatmap: %v", err)
	}
}
