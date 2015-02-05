package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"log"
	"os/exec"
	"strings"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//log.Println("event:", event)
				if (event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write) && strings.HasSuffix(event.Name, "go") {
					cmd := exec.Command("go", "test")
					out, _ := cmd.CombinedOutput()
					fmt.Printf("%s\n", out)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
