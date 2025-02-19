package main

import (
	"fmt"
)

var (
	binaries = map[string]string{
		"darwin_amd64":     "macos-x64",
		"darwin_arm64":     "macos-arm64",
		"linux_amd64":      "linux-x64",
		"linux_amd64_musl": "linux-x64-musl",
		"linux_arm64":      "linux-arm64",
		"linux_arm64_musl": "linux-arm64-musl",
		"linux_arm":        "linux-armv7",
		"windows_amd64":    "windows-x64",
		"windows_arm64":    "windows-arm64",
	}

	latestURL    = "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-%s"
	versionedURL = "https://github.com/tailwindlabs/tailwindcss/releases/download/%s/tailwindcss-%s"

	errUnsupportedArch = fmt.Errorf("unsupported os and architecture")
)

func buildURL(os, arch, version string, musl bool) (string, error) {
	selector := fmt.Sprintf("%v_%v", os, arch)
	if musl {
		selector += "_musl"
	}

	binary, ok := binaries[selector]
	if !ok {
		return "", fmt.Errorf("unsupported architecture %s: %w", selector, errUnsupportedArch)
	}

	// unspeficied version uses latest
	if version == "" {
		return fmt.Sprintf(latestURL, binary), nil
	}

	return fmt.Sprintf(versionedURL, version, binary), nil
}
