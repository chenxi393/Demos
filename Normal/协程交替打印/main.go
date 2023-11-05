package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup 要会用这个

func main() {
	PrintForSelect(1)
	// 如果有定数 wg.Wait() 这个也可以
	// 看看能不能关闭channel 来告知结束
	// 加锁实现应该是不合适的？？
	// PrintForSelect(1)
}

// select实现
func PrintForSelect(i int) {
	c1 := make(chan struct{})
	c2 := make(chan struct{})

	printAB := func() {
		for {
			select {
			case <-c1:
				fmt.Println(i)
				i++
				c2 <- struct{}{}
			case <-c2:
				fmt.Println(i)
				i++
				c1 <- struct{}{}
			}
		}
	}

	// go func() {
	// 	c1 <- struct{}{}
	// }()

	go printAB()
	go printAB()
	c1 <- struct{}{}
	time.Sleep(4 * time.Microsecond)
}
