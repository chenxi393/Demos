package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// 向通道发送数据
	ch <- 1

	// 关闭通道
	// 且channel是可以被GC机制回收掉的，所以关闭通道不是必须操作的
	close(ch)

	// 使用 select 判断通道是否已关闭
	// for {
	// 	select {
	// 	case value, ok := <-ch:
	// 		if !ok {
	// 			// 通道已关闭
	// 			fmt.Println("Channel is closed")
	// 			break
	// 		}
	// 		// 处理接收到的数据
	// 		fmt.Println("Received:", value)
	// 	}
	// }

	ch = make(chan int, 2) // 向通道发送数据
	ch <- 1
	ch <- 2   // 关闭通道
	close(ch) // 安全地从通道读取数据
	val, ok := <-ch
	fmt.Println(val, ok) // 输出: 1 true
	val, ok = <-ch
	fmt.Println(val, ok) // 输出: 2 true
	// 当通道中没有数据时，从已关闭的通道读取
	val, ok = <-ch
	fmt.Println(val, ok) // 输出: 0 false
	// 再次尝试向已关闭的通道发送数据将导致 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	ch <- 3 // 这将引发 panic

}
