package main

import (
	"fmt"
	"sync"
	"time"
)

func count(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup //等待队列
	wg.Add(2)
	go func() { //匿名函数
		count("🐏", 5)
		wg.Done() //等待队列减1
	}()
	go func() {
		count("🐂", 5)
		wg.Done()
	}()
	wg.Wait() //等待队列为0则返回 否则阻塞（or别的）

}
