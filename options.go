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

type Option func()

// UseConfigPath sets the path to the TailwindCSS configuration file
// otherwise it defaults to "config/tailwind.config.js".
func UseConfigPath(path string) Option {
	return func() {
		configPath = path
	}
}

// UseInputPath sets the path to the input file, the one with @apply rules
// otherwise it defaults to "web/assets/application.css".
func UseInputPath(path string) Option {
	return func() {
		inputPath = path
	}
}

// UseOutputPath sets the path to the output file, the one with compiled CSS
// otherwise it defaults to "web/public/application.css".
func UseOutputPath(path string) Option {
	return func() {
		outputPath = path
	}
}

// UseBinaryPath sets the path to the TailwindCSS CLI binary otherwise it
// defaults to "bin/tailwindcss".
func UseBinaryPath(path string) Option {
	return func() {
		binaryPath = path
	}
}
