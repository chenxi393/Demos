package main

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Println("接收到数据:", value)

	}
}

func testGoroutine() {
	var nums = []int{1, 2, 3}
	var wg sync.WaitGroup
	for i := range nums {
		wg.Add(1)
		go func() {
			fmt.Print(nums[i])
			wg.Done()
		}()
	}
	wg.Wait()
}

// 这才是正确的
func testGoroutineFixed() {
	var nums = []int{1, 2, 3}
	var wg sync.WaitGroup
	for i := range nums {
		wg.Add(1)
		go func(i int) {
			fmt.Print(nums[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func testDefer() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()
	i++
}

var b uint64

func testAddUint64() {
	for i := 0; i < 10000; i++ {
		// 原子加法
		go atomic.AddUint64(&b, uint64(1))
		// go func() {
		// 	b = b + 1
		// }()
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println(b)

}

func testReflect() {
	var x int = 8
	v := reflect.ValueOf(x)
	v.SetInt(24)
	println(x)
}

// 上面传入的x只是一个副本
func testReflectFixed() {
	var x int = 8
	v := reflect.ValueOf(&x).Elem()
	v.SetInt(24)
	println(x)
}

// 已关闭的channel 可以被读取 读完之后在读都是空值
var ch = make(chan int, 2)

func testChannel() {
	ch <- 3
	ch <- 4
	ch <- 9
	close(ch)
}

// 判断一个channel已经关闭
func IsClosed() {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 1; i <= 5; i++ {
			c <- i
		}
	}()

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Println("接收到数据:", value)
	}
}
