package main

import "fmt"

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := increment()
	// increment()是counter的父函数
	// 因此increment不会被GC 内的局部变量仍然有效
	fmt.Println(counter()) // 输出: 1
	fmt.Println(counter()) // 输出: 2
	fmt.Println(counter()) // 输出: 3
}
