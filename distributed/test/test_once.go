package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func setupConfig() {
	fmt.Println("Initializing configuration...")
	// 在这里执行一些只需执行一次的初始化操作
}

func main() {
	// 以下代码片段会调用 setupConfig() 函数，并确保只执行一次
	once.Do(setupConfig)

	// 这样不会调用
	once.Do(setupConfig)

	// 但是这样还是会调用函数
	setupConfig()
}
