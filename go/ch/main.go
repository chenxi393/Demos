package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// 向通道发送数据
	ch <- 1

	// 关闭通道
	close(ch)

	// 使用 select 判断通道是否已关闭
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				// 通道已关闭
				fmt.Println("Channel is closed")
				return
			}
			// 处理接收到的数据
			fmt.Println("Received:", value)
		}
	}
}
