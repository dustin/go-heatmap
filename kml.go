package heatmap

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
)

const kml_start = `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2">
<Folder>`
const kml_overlay = `
    <GroundOverlay>
      <Icon>
        <href>%s</href>
      </Icon>
      <LatLonBox>
        <north>%2.16f</north>
        <south>%2.16f</south>
        <east>%2.16f</east>
        <west>%2.16f</west>
        <rotation>0</rotation>
      </LatLonBox>
    </GroundOverlay>`
const kml_end = `</Folder></kml>`

const kml = kml_start + kml_overlay + kml_end

// Generate a heatmap for geographical data and the kml required to render it.
func HeatmapKML(size image.Rectangle, points []DataPoint, dotSize int, opacity uint8,
	scheme []color.Color, imgurl string, out io.Writer) (image.Image, error) {

	limits := findLimits(points)
	if limits.Min.X() < -180 || limits.Max.X() > 180 {
		return nil, errors.New("X limits out of range")
	}

	if limits.Min.Y() < -90 || limits.Max.Y() > 90 {
		return nil, errors.New("Y limits out of range")
	}

	mapimg := Heatmap(size, points, dotSize, opacity, scheme)

	adjustedLimits := adjustLimits(limits, size, dotSize)

	_, err := fmt.Fprintf(out, kml, imgurl,
		adjustedLimits.Max.Y(),
		adjustedLimits.Min.Y(),
		adjustedLimits.Max.X(),
		adjustedLimits.Min.X())

	return mapimg, err
}

func adjustLimits(limits limits, size image.Rectangle, dotSize int) (rv limits) {
	halfdot := float64(dotSize) / 2.0
	offx := (halfdot / float64(size.Dx()-dotSize)) * float64(limits.Dx())
	offy := (halfdot / float64(size.Dy()-dotSize)) * float64(limits.Dy())

	rv.Min = P(limits.Min.X()-offx, limits.Min.Y()-offy)
	rv.Max = P(limits.Max.X()+offx, limits.Max.Y()+offy)
	return
}
