package main

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/dustin/go-heatmap/schemes"
)

func main() {
	// This is what generated the prototype image for AlphaFire.
	spec := schemes.SchemeSpec{
		// From white to yellow
		schemes.SchemeRange{
			From:  color.White,
			To:    color.NRGBA{255, 255, 0, 255},
			Steps: 50,
		},
		// Then from yellow to red
		schemes.SchemeRange{
			From:  color.NRGBA{255, 255, 0, 255},
			To:    color.NRGBA{255, 0, 0, 255},
			Steps: 60,
		},
		// Then go from red to a bit of grey.
		schemes.SchemeRange{
			From:  color.NRGBA{255, 0, 0, 255},
			To:    color.NRGBA{128, 128, 128, 220},
			Steps: 100,
		},
		// Then down to transparent
		schemes.SchemeRange{
			From:  color.NRGBA{128, 128, 128, 220},
			To:    color.Transparent,
			Steps: 46,
		},
	}

	f, err := os.Create("test.png")
	if err != nil {
		log.Fatalf("Error making test.png: %v", err)
	}
	defer f.Close()
	png.Encode(f, spec)
}
