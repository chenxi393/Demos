package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var query = "test"

var matches int
var max_workers = 100
var now_workers int

var HasDone = make(chan bool)
var request = make(chan string)
var HasFound = make(chan bool)

func main() {
	start := time.Now()
	go search("/home/chenxi_393/",true)

	//TO DO  wait
	wait()

	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func wait() {
	for {
		select {
		case <-HasDone:
			now_workers--
			if now_workers == 0 {
				return
			}
		case path := <-request:
			go search(path, true)
			now_workers++
		case <-HasFound:
			matches++
		}
	}
}
func search(path string, isGo bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				HasFound <- true
			}
			if file.IsDir() {
				if now_workers < max_workers {
					request <- (path + file.Name()+"/")
				} else {
					search(path+file.Name()+"/", false)
				}
			}
		}
	}
	if isGo == true {
		HasDone <- true
	}
}
