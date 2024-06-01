package main

import "fmt"

// CSP并发模型 CSP讲究的是“以通信的方式来共享内存”。
var finish = make(chan struct{})

// 打印奇数
func printOdd(i int, c chan int) {
	for i < 1000 {
		fmt.Println(i)
		i++
		c <- i
		i = <-c
	}
	finish <- struct{}{}
}

// 打印偶数
func printEven(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
		i++
		c <- i
	}
}

func v1() {
	c := make(chan int)
	go printOdd(1, c)
	go printEven(c)
	<-finish
}
