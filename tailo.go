// tailo is a wrapper for the Tailwind CSS CLI that
// facilitates the download and of the CLI and the
// config file.
package tailo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type extensions []string

func (e extensions) Has(ext string) bool {
	for _, v := range e {
		if v == ext {
			return true
		}
	}

	return false
}

// Extensions to watch for changes.
var watchExtensions = extensions{".html", ".css"}

func buildWatcher() (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return watcher, fmt.Errorf("error creating watcher: %w", err)
	}

	// Add all files that need to be watched to the
	// watcher so it notifies the errors that it needs to
	// restart.
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !watchExtensions.Has(filepath.Ext(path)) {
			return nil
		}

		return watcher.Add(path)
	})

	if err != nil {
		return watcher, fmt.Errorf("error loading paths: %w", err)
	}

	return watcher, err
}
