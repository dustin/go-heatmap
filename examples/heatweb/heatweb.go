package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
)

const maxInputLength = 10000

type csvpoint []string

func (c csvpoint) X() float64 {
	x, _ := strconv.ParseFloat(c[0], 64)
	return x
}

func (c csvpoint) Y() float64 {
	x, _ := strconv.ParseFloat(c[1], 64)
	return x
}

func parseInt(vals url.Values, v string, def, min, max int) int {
	rv, err := strconv.ParseInt(vals.Get(v), 10, 32)
	if err != nil || int(rv) < min || int(rv) > max {
		return def
	}
	return int(rv)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {

	vals := req.URL.Query()

	width := parseInt(vals, "w", 1024, 100, 4096)
	height := parseInt(vals, "h", 768, 100, 4096)
	dotsize := parseInt(vals, "d", 200, 1, 256)
	opacity := uint8(parseInt(vals, "o", 128, 1, 255))

	defer req.Body.Close()
	lr := io.LimitedReader{req.Body, maxInputLength}
	cr := csv.NewReader(&lr)

	data := []heatmap.DataPoint{}
	reading := true
	for reading {
		rec, err := cr.Read()
		switch err {
		case io.EOF:
			reading = false
		case nil:
			data = append(data, csvpoint(rec))
		default:
			log.Printf("Other error:  %#v", err)
			w.WriteHeader(400)
			fmt.Fprintf(w, "Error reading data: %v", err)
			return
		}
	}

	w.Header().Set("Content-type", "application/vnd.google-earth.kmz")
	w.WriteHeader(200)
	heatmap.HeatmapKMZ(image.Rect(0, 0, width, height),
		data, dotsize, opacity, schemes.AlphaFire, w)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":1756", nil))
}
