package heatmap

import (
	"archive/zip"
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/dustin/go-heatmap/schemes"
	"github.com/jteeuwen/imghash"
)

const testKmlImgURL = "http://www.example.com/thing.png"

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

const expKmz = `<?xml version="1.0" encoding="UTF-8"?>
		<kml xmlns="http://www.opengis.net/kml/2.2">
		<Folder>
		    <GroundOverlay>
		      <Icon>
		        <href>heatmap.png</href>
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

func TestKML(t *testing.T) {
	kmlBuf := &bytes.Buffer{}

	img, err := KML(image.Rect(0, 0, 1024, 1024),
		testPoints, 150, 128, schemes.AlphaFire,
		testKmlImgURL, kmlBuf)
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

func TestKMLOutOfRange(t *testing.T) {
	kmlBuf := &bytes.Buffer{}

	_, err := KML(image.Rect(0, 0, 1024, 1024),
		append(testPoints, P(-200, 0)), 150, 128, schemes.AlphaFire,
		testKmlImgURL, kmlBuf)
	if err == nil {
		t.Fatalf("Expected error with bad input")
	}
}

func TestKMZBadInput(t *testing.T) {
	err := KMZ(image.Rect(0, 0, 1024, 1024),
		append(testPoints, P(-200, 0)), 150, 128, schemes.AlphaFire,
		ioutil.Discard)
	if err == nil {
		t.Fatalf("Expected error with bad input")
	}
}

type writeFailer struct {
	n int
}

func (w *writeFailer) Write(p []byte) (int, error) {
	towrite := len(p)
	var err error
	if towrite > w.n {
		towrite = w.n
		err = io.EOF
	}
	w.n -= towrite
	return towrite, err
}

func TestKMZBadWriter(t *testing.T) {
	err := KMZ(image.Rect(0, 0, 1024, 1024),
		testPoints, 150, 128, schemes.AlphaFire,
		&writeFailer{514})
	if err == nil {
		t.Fatalf("Expected error with bad input")
	}
}

func rzd(t *testing.T, zf *zip.File) []byte {
	r, err := zf.Open()
	if err != nil {
		t.Fatalf("Error opening %v: %v", zf.Name, err)
	}
	defer r.Close()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("Error reading %v: %v", zf.Name, err)
	}
	return data
}

func TestKMZ(t *testing.T) {
	buf := &bytes.Buffer{}
	err := KMZ(image.Rect(0, 0, 1024, 1024),
		testPoints, 150, 128, schemes.AlphaFire,
		buf)
	if err != nil {
		t.Fatalf("Failed to build kmz: %v", err)
	}
	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		t.Fatalf("Error reading kmz: %v", err)
	}

	for _, zf := range zr.File {
		switch zf.Name {
		case "doc.kml":
			gotK := rzd(t, zf)
			if !xsimilar(string(gotK), expKmz) {
				t.Errorf("Expected kml=%v, got %s", expKmz, gotK)
			}
		case "heatmap.png":
			img, _, err := image.Decode(bytes.NewReader(rzd(t, zf)))
			if err != nil {
				t.Errorf("Error decoding image: %v", err)
			}
			got := imghash.Average(img)
			if got != expHash {
				t.Errorf("Expected image hash %v, got %v", expHash, got)
			}
		default:
			t.Fatalf("Unexpected zipfile name: %v", err)
		}
	}
}
