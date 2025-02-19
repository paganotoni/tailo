package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	version string
	musl    bool

	// the place where the binary will be downloaded
	// and used to execute the Tailwind CSS command
	binary = "bin/tailwindcss"
)

func main() {
	args := os.Args

	fs := flag.NewFlagSet("download", flag.ExitOnError)
	fs.BoolVarP(&musl, "musl", "m", false, "download musl build for linux")
	fs.StringVarP(&version, "version", "v", "", "Version of Tailwind CSS to download")

	if len(args) == 1 {
		fmt.Println("Usage:")
		fmt.Println("  tailo download [-v <version>] [--musl]")
		fmt.Println("  tailo [--input input.css] [--output output.css] [--watch] [options…]")

		fmt.Print("\nFlags:\n")
		fs.PrintDefaults()

		return
	}

	if args[1] == "download" {
		fs.Parse(os.Args[1:])

		err := download(binary, version, musl)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		return
	}

	// Run tailwind command with the rest of args
	err := run(binary, args[1:])
	if err != nil {
		os.Exit(1)
	}
}
