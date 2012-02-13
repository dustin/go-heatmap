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
	// This is what generated the prototype image for AlphaFire.
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
		// Then go from red to a bit of grey.
		schemes.SchemeRange{
			color.NRGBA{255, 0, 0, 255},
			color.NRGBA{128, 128, 128, 220},
			100,
		},
		// Then down to transparent
		schemes.SchemeRange{
			color.NRGBA{128, 128, 128, 220},
			color.Transparent,
			46,
		},
	}

	colors := schemes.Build(spec)

	i := image.NewRGBA(image.Rect(0, 0, 64, len(colors)))
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
