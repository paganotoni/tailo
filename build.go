package tailo

import (
	"fmt"
	"os"
	"os/exec"
)

// Build runs the Tailwind CSS CLI binary to build the
// CSS file and generate compiled CSS it expects to find
// the options in the config file.
func Build(options ...Option) {
	// Applying passed options
	for _, option := range options {
		option()
	}

	if err := Setup(); err != nil {
		fmt.Println("Error running the setup:", err.Error())
		os.Exit(1)

		return
	}

	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		err := Setup()
		if err != nil {
			panic(err)
		}
	}

	cmd := exec.Command(binaryPath)
	cmd.Args = append(cmd.Args, "--config", configPath)
	cmd.Args = append(cmd.Args, "--input", inputPath)
	cmd.Args = append(cmd.Args, "--output", outputPath)

	cmd.Stdout = writer(0)
	cmd.Stderr = writer(0)

	fmt.Println("[tailo] Running:", cmd.String())

	if err := cmd.Run(); err != nil {
		fmt.Println("[tailo] Error running tailwindcss:", err)
	}
}
