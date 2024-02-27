// tailo is a wrapper for the Tailwind CSS CLI that
// facilitates the download and of the CLI and the
// config file.
package tailo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

var (
	binaries = map[string]string{
		"darwin_amd64":  "tailwindcss-macos-x64",
		"darwin_arm64":  "tailwindcss-macos-arm64",
		"linux_amd64":   "tailwindcss-linux-x64",
		"linux_arm64":   "tailwindcss-linux-arm64",
		"linux_arm":     "tailwindcss-linux-armv7",
		"windows_amd64": "tailwindcss-windows-x64",
		"windows_arm64": "tailwindcss-windows-arm64",
	}
)

// Setup downloads the Tailwind CSS CLI binary for the
// given operating system and architecture. It makes the
// binary executable and places it in the bin/ directory.
func Setup() error {
	if _, err := os.Stat("bin/tailwindcss"); err == nil {
		fmt.Println("Tailwind CSS CLI binary already exists.")

		return nil
	}

	currentArch := fmt.Sprintf("%v_%v", runtime.GOOS, runtime.GOARCH)
	binary, ok := binaries[currentArch]
	if !ok {
		return fmt.Errorf("Unsupported operating system and architecture: %s", currentArch)
	}

	url := fmt.Sprintf("https://github.com/tailwindlabs/tailwindcss/releases/latest/download/%v", binary)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Could not download Tailwind CSS CLI binary: %s", resp.Status)
	}

	err = os.MkdirAll("bin", 0755)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create("bin/tailwindcss")
	if err != nil {
		return err
	}

	defer out.Close()

	err = os.Chmod("bin/tailwindcss", 0755)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	return err
}
