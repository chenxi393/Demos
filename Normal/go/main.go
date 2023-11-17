package main

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
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
