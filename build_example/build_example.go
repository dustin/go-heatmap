package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/dustin/go-heatmap/schemes"
)

func main() {
	// This is somewhat, but not exactly similar to alphafire.
	spec := []schemes.SchemeRange{
		// From white to yellow
		schemes.SchemeRange{
			color.White,
			color.NRGBA{255, 255, 0, 255},
			50,
		},
		// Then from yellow to red
		schemes.SchemeRange{
			color.NRGBA{255, 255, 0, 255},
			color.NRGBA{255, 0, 0, 255},
			60,
		},
		// Then fade red down to transparent
		schemes.SchemeRange{
			color.NRGBA{255, 0, 0, 255},
			color.Transparent,
			150,
		},
	}

	colors := schemes.Build(spec)

	i := image.NewRGBA(image.Rect(0, 0, 256, len(colors)))
	for y, c := range colors {
		for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
			i.Set(x, y, c)
		}
	}

	f, err := os.Create("test.png")
	if err != nil {
		log.Fatalf("Error making test.png: %v", err)
	}
	defer f.Close()
	png.Encode(f, i)
}
