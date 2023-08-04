package tailo

var (
	// configPath is the path to the TailwindCSS configuration file.
	configPath = "config/tailwind.config.js"

	// inputPath is the path to the input file, the one with @apply rules.
	inputPath = "web/assets/application.css"

	// outputPath is the path to the output file, the one with compiled CSS.
	outputPath = "web/public/application.css"

	// binaryPath is the path to the TailwindCSS CLI binary.
	binaryPath = "bin/tailwindcss"
)

type option func()

// UseConfigPath sets the path to the TailwindCSS configuration file
// otherwise it defaults to "config/tailwind.config.js".
func UseConfigPath(path string) option {
	return func() {
		configPath = path
	}
}

// UseInputPath sets the path to the input file, the one with @apply rules
// otherwise it defaults to "web/assets/application.css".
func UseInputPath(path string) option {
	return func() {
		inputPath = path
	}
}

// UseOutputPath sets the path to the output file, the one with compiled CSS
// otherwise it defaults to "web/public/application.css".
func UseOutputPath(path string) option {
	return func() {
		outputPath = path
	}
}

// UseBinaryPath sets the path to the TailwindCSS CLI binary otherwise it
// defaults to "bin/tailwindcss".
func UseBinaryPath(path string) option {
	return func() {
		binaryPath = path
	}
}
