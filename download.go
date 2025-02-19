// tailo is a wrapper for the Tailwind CSS CLI that
// facilitates the download and of the CLI and the
// config file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// download the Tailwind CSS CLI binary for the
// given operating system and architecture. It makes the
// binary executable and places it in the bin/ directory.
func download(binary, version string, musl bool) error {
	if _, err := os.Stat(binary); err == nil {
		err := os.Remove(binary)
		if err != nil {
			return err
		}
	}

	// Build the URL for the binary
	url, err := buildURL(runtime.GOOS, runtime.GOARCH, version, musl)
	if err != nil {
		return err
	}

	fmt.Println("Downloading:", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("error downloading binary, got %s", resp.Status)
	}

	err = os.MkdirAll(filepath.Dir(binary), 0o755)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(binary)
	if err != nil {
		return err
	}

	defer out.Close()

	err = os.Chmod(binary, 0o755)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("CLI downloaded successfully to: %s\n", binary)
	return nil
}
