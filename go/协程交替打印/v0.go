package main

import "fmt"

var i = 1
var stop = 10000000

// 打印奇数
func printOdd0(c chan struct{}) {
	for {
		fmt.Println(i)
		i++
		//time.Sleep(1 * time.Second)
		c <- struct{}{}
		<-c
	}
}

// 打印偶数
func printEven0(c chan struct{}) {
	for {
		<-c
		fmt.Println(i)
		i++
		//time.Sleep(1 * time.Second)
		c <- struct{}{}
	}
}
