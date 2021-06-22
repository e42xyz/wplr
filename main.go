package main 

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sync"

    "github.com/jaschaephraim/lrserver"
    "github.com/fsnotify/fsnotify"
)

var (
	watchDirPath = flag.String("p", "/var/www/html", "The path to the WordPress install.")
	watcher *fsnotify.Watcher
)

func Watcher() {
    var wg sync.WaitGroup
    // Create file watcher
    watcher, _ = fsnotify.NewWatcher()
    defer watcher.Close()

    // Create and start LiveReload server
    lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
    go lr.ListenAndServe()

    if err := filepath.Walk(*watchDirPath, watchDir); err != nil {
        log.Println("ERROR", err)
    }

    //
    done := make(chan bool)
    wg.Add(1)
    // Start goroutine that requests reload upon watcher event
    go func() {
        for {
            select {
            case event := <-watcher.Events:
                lr.Reload(event.Name)
            case err := <-watcher.Errors:
                log.Println(err)
            }
        }
        defer wg.Done()
    }()

    <-done

    wg.Wait()
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {

    // since fsnotify can watch all the files in a directory, watchers only need
    // to be added to each nested directory
    if fi.Mode().IsDir() {
        return watcher.Add(path)
    }

    return nil
}

func wplrUsage() {
    fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
    flag.PrintDefaults()
}

func main() {
	// TODO: switch to config file
    flag.Usage = wplrUsage
    flag.Parse()
    // if WordPress path isn't provided print usage and exit
	if len(os.Args) < 2 {
        wplrUsage()
        os.Exit(0)
    }
	// Start the watcher
    Watcher()
}