package main

import (
	"errors"
	"testing"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name    string
		goos    string
		goarch  string
		musl    bool
		version string
		want    string
		err     error
	}{
		{
			name:    "darwin amd64 latest",
			goos:    "darwin",
			goarch:  "amd64",
			musl:    false,
			version: "",
			want:    "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64",
		},
		{
			name:    "linux arm64 with version",
			goos:    "linux",
			goarch:  "arm64",
			musl:    false,
			version: "v3.0.0",
			want:    "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.0.0/tailwindcss-linux-arm64",
		},
		{
			name:    "linux amd64 with musl",
			goos:    "linux",
			goarch:  "amd64",
			musl:    true,
			version: "",
			want:    "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64-musl",
		},
		{
			name:    "windows amd64",
			goos:    "windows",
			goarch:  "amd64",
			musl:    false,
			version: "",
			want:    "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-windows-x64",
		},
		{
			name:    "unsupported platform",
			goos:    "solaris",
			goarch:  "amd64",
			musl:    false,
			version: "",
			err:     errUnsupportedArch,
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildURL(tt.goos, tt.goarch, tt.version, tt.musl)
			if tt.err != nil && !errors.Is(err, tt.err) {
				t.Errorf("buildURL() error = %v, wantErr %v", err, tt.err)
			}

			if got != tt.want {
				t.Errorf("buildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
