package filesystem

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

func WatchDirectories() *fsnotify.Watcher {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)

				if event.Has(fsnotify.Create) {
					go func() {
						time.Sleep(time.Second * 15)

						_, err := os.Stat(event.Name)
						if err != nil {
							log.Println("file already deleted")
							return
						}

						os.Remove(event.Name)
					}()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add both folders to watch.
	err = watcher.Add(fmt.Sprintf("%s/assets/posts/", dir))
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(fmt.Sprintf("%s/assets/users", dir))
	if err != nil {
		log.Fatal(err)
	}

	return watcher
}
