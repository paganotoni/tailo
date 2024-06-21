package main

import (
	"flag"

	"github.com/paganotoni/tailo"
)

var (
	input  string
	output string
	config string
	binary string
)

func init() {
	flag.StringVar(&input, "input", "internal/assets/application.css", "The CSS input file path.")
	flag.StringVar(&output, "output", "public/application.css", "The CSS output file path.")
	flag.StringVar(&config, "config", "tailwind.config.js", "The TailwindCSS configuration file path.")
	flag.StringVar(&binary, "binary", "bin/tailwindcss", "The TailwindCSS CLI binary path.")
}

func main() {
	flag.Parse()

	tailo.Setup()
	tailo.Build(
		tailo.UseInputPath(input),
		tailo.UseOutputPath(output),
		tailo.UseConfigPath(config),
		tailo.UseBinaryPath(binary),
	)
}
