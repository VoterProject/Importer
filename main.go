package main

import (
	"flag"
	"github.com/voterproject/importer/app/importer"
	"os"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		panic("Please provide data path")
	}

	path := flag.Arg(0)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	importer.Start(path)
}
