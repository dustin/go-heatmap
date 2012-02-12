package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/dustin/go-heatmap/schemes"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatalf("What are you wanting to compile?")
	}
	scheme, err := schemes.FromImage(os.Args[1])
	if err != nil {
		log.Fatalf("Error initializing scheme: %v", err)
	}

	fmt.Printf("package whatever\n\n")
	fmt.Printf("import \"image/color\"\n\n")
	fmt.Printf("var MyAwesomeScheme = []color.Color{\n")
	for _, cin := range scheme {
		c := cin.(color.RGBA)
		fmt.Printf("\tcolor.RGBA{%v, %v, %v, %v},\n",
			c.R, c.B, c.G, c.A)
	}
	fmt.Printf("}")
}
