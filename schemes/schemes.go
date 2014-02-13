// Package schemes provides standard color schemes from gheat.
//
// These color schemes were converted from the pngs from gheat are made
// available for your convenience.
package schemes

import (
	"image"
	"image/color"
	_ "image/gif"  // image format side effect
	_ "image/jpeg" // image format side effect
	_ "image/png"  // image format side effect
	"os"
)

// FromImage creates a color scheme from an image.
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

type mycolor struct {
	R float64
	G float64
	B float64
	A float64
}

func (c mycolor) RGBA() (R, G, B, A uint32) {
	// My colors aren't pre-multiplied, so I've got to convert
	// myself to an NRGBA, then an RGBA, then expand.

	return color.RGBAModel.Convert(color.NRGBA{
		uint8(c.R / 256.0),
		uint8(c.G / 256.0),
		uint8(c.B / 256.0),
		uint8(c.A / 256.0),
	}).RGBA()
}

// Build generates a heatmap color scheme smoothly across the given
// scheme ranges.
func Build(spec SchemeSpec) []color.Color {
	rv := []color.Color{}

	for _, r := range spec {
		fr, fg, fb, fa := floatRGBA(r.From)
		tr, tg, tb, ta := floatRGBA(r.To)

		rd := (tr - fr) / float64(r.Steps)
		gd := (tg - fg) / float64(r.Steps)
		bd := (tb - fb) / float64(r.Steps)
		ad := (ta - fa) / float64(r.Steps)

		for x := 0; x < r.Steps; x++ {
			rv = append(rv, mycolor{fr, fg, fb, fa})

			fr += rd
			fg += gd
			fb += bd
			fa += ad
		}
	}

	return rv
}

func floatRGBA(c color.Color) (r, g, b, a float64) {
	ir, ig, ib, ia := color.NRGBAModel.Convert(c).RGBA()
	return float64(ir), float64(ig), float64(ib), float64(ia)
}

// A SchemeRange is a color range for computing a scheme.
type SchemeRange struct {
	// The starting color for this range
	From color.Color
	// The ending color for this range
	To color.Color
	// The number of colors to which this range should expand
	Steps int
}

// A SchemeSpec is a color generator specification.
//
// Basically, this is just a bunch of start/stop points and colors
// along a range.
type SchemeSpec []SchemeRange

// ColorModel satisfies the image.Image interface.
func (ss SchemeSpec) ColorModel() color.Model {
	return color.NRGBAModel
}

// Bounds satisfies the image.Image interface.
func (ss SchemeSpec) Bounds() image.Rectangle {
	return image.Rect(0, 0, 32, 256)
}

// At satisfies the image.Image interface.
func (ss SchemeSpec) At(x, y int) color.Color {
	// This is the ridiculously slow version.
	return Build(ss)[y]
}
