package main

import (
	"bufio"
	"encoding/json"
	"image"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
	"github.com/dustin/go-humanize"
)

const url = "http://single.couchbase.net/wikipedia2012/_design/places/_spatial/points?bbox=-180,-90,180,90"

type row struct {
	ID       string `json:"id"`
	Value    int    `json:"value"`
	Geometry struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

func main() {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error opening strem from %q: %v", url, err)
	}
	defer res.Body.Close()

	r := bufio.NewReader(res.Body)
	_, _, err = r.ReadLine()
	if err != nil {
		log.Fatalf("Error reading first line: %v", err)
	}

	locations := make([]heatmap.DataPoint, 0, 200000)

	for {
		bytes, isPrefix, err := r.ReadLine()
		if err != nil {
			log.Fatalf("Error reading line: %v", err)
		}
		if isPrefix {
			log.Fatalf("Crap, that was a prefix...")
		}
		if bytes[0] != '{' {
			break
		}
		r := row{}
		if bytes[len(bytes)-1] == ',' {
			bytes = bytes[:len(bytes)-1]
		}
		err = json.Unmarshal(bytes, &r)
		if err != nil {
			log.Printf("Couldn't parse %v: %v", string(bytes), err)
			break
		}

		locations = append(locations, heatmap.P(r.Geometry.Coordinates[0],
			r.Geometry.Coordinates[1]))
	}
	end := time.Now()

	log.Printf("Parsed %s items in %s",
		humanize.Comma(int64(len(locations))), end.Sub(start))

	out, err := os.Create("wikipedia.kmz")
	if err != nil {
		log.Fatalf("Error making output file:  %v", err)
	}
	defer out.Close()

	err = heatmap.HeatmapKMZ(image.Rect(0, 0, 8192, 4096), locations, 50, 96,
		schemes.AlphaFire, out)
	if err != nil {
		log.Fatalf("Error generating thingy: %v", err)
	}

	log.Printf("Completed heatmap generation in %s", time.Now().Sub(end))
}
