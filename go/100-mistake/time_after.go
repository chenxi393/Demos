package main

import (
	"fmt"
	"log"
	"time"
	"unsafe"
)

// 这个也todo吧 感觉并不常用
// time.After导致内存泄露
func consumer(ch <-chan string) {
	for {
		select {
		case event := <-ch:
			// handle(event)
			fmt.Println(event)
			time.Sleep(2 * time.Second)
		case <-time.After(time.Hour):
			log.Println("warning: no messages received")
		}
	}
}

func consumer3(ch <-chan string) {
	timerDuration := 1 * time.Hour
	// time.Duration实际上是int64类型的别名，它的单位是纳秒。这里传的是1000纳秒，也就是1微秒
	timer := time.NewTimer(timerDuration)
	defer timer.Stop()
	for {
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			// handle(event)
			fmt.Println(event)
			time.Sleep(2 * time.Second)
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}

func test() {
	// consumer(make(<-chan string))

	c := make(chan time.Time, 1)

	fmt.Println(unsafe.Sizeof(c))
}
