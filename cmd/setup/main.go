package main

import (
	"flag"

	"github.com/paganotoni/tailo"
)

var (
	input   string
	output  string
	config  string
	binary  string
	version string
)

func init() {
	flag.StringVar(&version, "version", "", "The TailwindCSS version to use, defaults to empty which means latest.")
}

func main() {
	flag.Parse()

	tailo.Setup(
		tailo.UseVersion(version),
	)
}
