package tailo

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func Watch(options ...option) {
	Build(options...)

	watcher, err := buildWatcher()
	if err != nil {
		panic(err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					Build(options...)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-make(chan struct{})
}
