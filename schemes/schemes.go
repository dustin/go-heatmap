/*
 Standard color schemes from gheat.

 These color schemes were converted from the pngs from gheat are made
 available for your convenience.
*/
package schemes

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// Create a color scheme from an image.
// The image should have all of the colors you want to use arranged vertically
// with the "coldest" on the bottom and the "hottest" on the top.
func FromImage(filename string) ([]color.Color, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	rv := []color.Color{}
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rv = append(rv, img.At(0, y))
	}
	return rv, nil
}

// Generate a heatmap color scheme smoothly across the given scheme ranges.
func Build(spec []SchemeRange) []color.Color {
	rv := []color.Color{}

	for _, r := range spec {
		fr, fg, fb, fa := floatRGBA(r.From)
		tr, tg, tb, ta := floatRGBA(r.To)

		rd := (tr - fr) / float64(r.Steps)
		gd := (tg - fg) / float64(r.Steps)
		bd := (tb - fb) / float64(r.Steps)
		ad := (ta - fa) / float64(r.Steps)

		for x := 0; x < r.Steps; x++ {
			c := color.RGBA{
				uint8(int(fr) / 256),
				uint8(int(fg) / 256),
				uint8(int(fb) / 256),
				uint8(int(fa) / 256),
			}
			fr += rd
			fg += gd
			fb += bd
			fa += ad

			rv = append(rv, c)
		}
	}

	return rv
}

func floatRGBA(c color.Color) (r, g, b, a float64) {
	ir, ig, ib, ia := c.RGBA()
	return float64(ir), float64(ig), float64(ib), float64(ia)
}

// A color scheme range for computing a scheme.
type SchemeRange struct {
	// The starting color for this range
	From color.Color
	// The ending color for this range
	To color.Color
	// The number of colors to which this range should expand
	Steps int
}
