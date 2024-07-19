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
	flag.StringVar(&input, "input", "internal/assets/application.css", "The CSS input file path.")
	flag.StringVar(&output, "output", "public/application.css", "The CSS output file path.")
	flag.StringVar(&config, "config", "tailwind.config.js", "The TailwindCSS configuration file path.")
	flag.StringVar(&binary, "binary", "bin/tailwindcss", "The TailwindCSS CLI binary path.")
	flag.StringVar(&version, "version", "", "The TailwindCSS version to use, defaults to empty which means latest.")
}

func main() {
	flag.Parse()

	tailo.Build(
		tailo.UseVersion(version),
		tailo.UseInputPath(input),
		tailo.UseOutputPath(output),
		tailo.UseConfigPath(config),
		tailo.UseBinaryPath(binary),
	)
}
