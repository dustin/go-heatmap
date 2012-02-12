package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dustin/go-heatmap/schemes"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: schemecompiler [-flags] infile.png\n")
		flag.PrintDefaults()
	}
}

func main() {

	var pkg = flag.String("package", "example", "scheme's package")
	var name = flag.String("name", "MyAwesomeScheme", "scheme's name")

	flag.Parse()
	log.SetFlags(0)
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	scheme, err := schemes.FromImage(flag.Arg(0))
	if err != nil {
		log.Fatalf("Error initializing scheme: %v", err)
	}

	fmt.Printf("package %s\n\n", *pkg)
	fmt.Printf("import \"image/color\"\n\n")
	fmt.Printf("var %s []color.Color\n\n", *name)

	fmt.Printf("func init() {\n")
	fmt.Printf("\t%s = []color.Color{\n", *name)
	for _, c := range scheme {
		fmt.Printf("\t\t%#v,\n", c)
	}
	fmt.Printf("\t}\n}\n")
}
