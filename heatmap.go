// Generate heatmaps.
package heatmap

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// A data point to be plotted.
// These are all normalized to use the maximum amount of
// space available in the output image.
type DataPoint interface {
	X() float64
	Y() float64
}

type apoint struct {
	x float64
	y float64
}

func (a apoint) X() float64 {
	return a.x
}

func (a apoint) Y() float64 {
	return a.y
}

// Construct a simple datapoint
func P(x, y float64) DataPoint {
	return apoint{x, y}
}

type limits struct {
	Min DataPoint
	Max DataPoint
}

// Draw a heatmap.
//
// size is the size of the image to crate
// dotSize is the impact size of each point on the output
// opacity is the alpha value (0-255) of the impact of the image overlay
// scheme is the color palette to choose from the overlay
func Heatmap(size image.Rectangle, points []DataPoint, dotSize int, opacity uint8,
	scheme []color.Color) image.Image {

	dot := mkDot(float64(dotSize))

	limits := findLimits(points)

	// Draw black/alpha onto this
	bw := image.NewRGBA(size)
	for _, p := range points {
		limits.placePoint(p, bw, dot)
	}

	rv := image.NewRGBA(size)

	// Then we transplant the pixels one at a time pulling from our color map
	warm(rv, bw, opacity, scheme)
	return rv
}

func warm(out, in draw.Image, opacity uint8, colors []color.Color) {
	bounds := in.Bounds()
	collen := float64(len(colors))
	invisible := color.NRGBA{255, 255, 255, 0}
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			col := in.At(x, y)
			_, _, _, alpha := col.RGBA()
			percent := float64(alpha) / float64(0xffff)
			var outcol color.Color
			if percent == 0 {
				outcol = invisible
			} else {
				template := colors[int((collen-1)*(1.0-percent))]
				tr, tg, tb, _ := template.RGBA()
				outcol = color.NRGBA{
					uint8(tr / 256),
					uint8(tg / 256),
					uint8(tb / 256),
					uint8(opacity)}
			}
			out.Set(x, y, outcol)
		}
	}
}

func findLimits(points []DataPoint) limits {
	minx, miny := points[0].X(), points[0].Y()
	maxx, maxy := minx, miny

	for _, p := range points {
		minx = math.Min(p.X(), minx)
		miny = math.Min(p.Y(), miny)
		maxx = math.Max(p.X(), maxx)
		maxy = math.Max(p.Y(), maxy)
	}

	return limits{apoint{minx, miny}, apoint{maxx, maxy}}
}

func mkDot(size float64) draw.Image {
	i := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))

	md := 0.5 * math.Sqrt(math.Pow(float64(size)/2.0, 2)+math.Pow((float64(size)/2.0), 2))
	for x := float64(0); x < size; x++ {
		for y := float64(0); y < size; y++ {
			d := math.Sqrt(math.Pow(x-size/2.0, 2) + math.Pow(y-size/2.0, 2))
			if d < md {
				rgbVal := uint8(200.0*d/md + 50.0)
				rgba := color.NRGBA{0, 0, 0, 255 - rgbVal}
				i.Set(int(x), int(y), rgba)
			}
		}
	}

	return i
}

func (l limits) translate(p DataPoint, i draw.Image, dotsize int) (rv image.Point) {
	// Normalize to 0-1
	x := float64(p.X()-l.Min.X()) / float64(l.Max.X()-l.Min.X())
	y := float64(p.Y()-l.Min.Y()) / float64(l.Max.Y()-l.Min.Y())

	// And remap to the image
	rv.X = int(x * float64((i.Bounds().Max.X - dotsize)))
	rv.Y = int((1.0 - y) * float64((i.Bounds().Max.Y - dotsize)))

	return
}

func (l limits) placePoint(p DataPoint, i, dot draw.Image) {
	pos := l.translate(p, i, dot.Bounds().Max.X)
	dotw, doth := dot.Bounds().Max.X, dot.Bounds().Max.Y
	draw.Draw(i, image.Rect(pos.X, pos.Y, pos.X+dotw, pos.Y+doth), dot,
		image.ZP, draw.Over)
}
