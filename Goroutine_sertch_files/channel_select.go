package main

import (
	"fmt"
	"time"
)

//往channel收听和发送消息都会阻塞代码的运行

func main() {
	c1 := make(chan string) //channel间通信
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "🐂"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "🐏"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		select { //选择一个没有被阻塞的channel
		case message := <-c1:
			fmt.Println(message)

		case message := <-c2:
			fmt.Println(message)
		}

	}
}
