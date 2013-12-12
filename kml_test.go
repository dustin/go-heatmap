package heatmap

import (
	"bytes"
	"image"
	"strings"
	"testing"

	"github.com/dustin/go-heatmap/schemes"
	"github.com/jteeuwen/imghash"
)

const testKmlImgUrl = "http://www.example.com/thing.png"

const expKml = `<?xml version="1.0" encoding="UTF-8"?>
		<kml xmlns="http://www.opengis.net/kml/2.2">
		<Folder>
		    <GroundOverlay>
		      <Icon>
		        <href>http://www.example.com/thing.png</href>
		      </Icon>
		      <LatLonBox>
		        <north>1.0844535751966644</north>
		        <south>-0.0815174737483185</south>
		        <east>1.0688449455069537</east>
		        <west>-0.0780794272341640</west>
		        <rotation>0</rotation>
		      </LatLonBox>
		    </GroundOverlay></Folder></kml>`

func xsimilar(a, b string) bool {
	return strings.Join(strings.Fields(a), " ") ==
		strings.Join(strings.Fields(b), " ")
}

func TestHeatmapKML(t *testing.T) {
	kmlBuf := &bytes.Buffer{}

	img, err := HeatmapKML(image.Rect(0, 0, 1024, 1024),
		testPoints, 150, 128, schemes.AlphaFire,
		testKmlImgUrl, kmlBuf)
	if err != nil {
		t.Fatalf("Error generating kml: %v", err)
	}
	got := imghash.Average(img)
	if got != expHash {
		t.Errorf("Expected image hash %v, got %v", expHash, got)
	}
	gotK := kmlBuf.String()
	if !xsimilar(gotK, expKml) {
		t.Errorf("Expected kml=%v, got %v", expKml, gotK)
	}
}
