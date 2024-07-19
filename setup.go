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
	version  string
	binaries = map[string]string{
		"darwin_amd64":  "macos-x64",
		"darwin_arm64":  "macos-arm64",
		"linux_amd64":   "linux-x64",
		"linux_arm64":   "linux-arm64",
		"linux_arm":     "linux-armv7",
		"windows_amd64": "windows-x64",
		"windows_arm64": "windows-arm64",
	}

	latestURL    = "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-%s"
	versionedURL = "https://github.com/tailwindlabs/tailwindcss/releases/download/%s/tailwindcss-%s"
)

// Setup downloads the Tailwind CSS CLI binary for the
// given operating system and architecture. It makes the
// binary executable and places it in the bin/ directory.
func Setup(options ...Option) error {
	// Running options
	for _, option := range options {
		option()
	}

	if _, err := os.Stat("bin/tailwindcss"); err == nil {
		fmt.Println("Tailwind CSS CLI binary already exists.")

		return nil
	}

	currentArch := fmt.Sprintf("%v_%v", runtime.GOOS, runtime.GOARCH)
	binary, ok := binaries[currentArch]
	if !ok {
		return fmt.Errorf("Unsupported operating system and architecture: %s", currentArch)
	}

	url := latestURL
	args := []any{binary}

	if version != "" {
		url = versionedURL
		args = []any{version, binary}
	}

	url = fmt.Sprintf(url, args...)
	fmt.Println("Downloading from:", url)

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
