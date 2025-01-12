package main

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

/*
这个错误是因为在Go语言中，接口赋值的规则是：
如果一个类型实现了某个接口的所有方法，那么该类型的实例（值类型或指针类型）就可以赋值给该接口。
如果一个类型实现了某个接口的所有方法，但是这些方法的接收者是指针类型而不是该类型本身的值类型，那么只有该类型的指针类型才可以赋值给该接口。
*/
type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func main() {
	var (
		wg    = &sync.WaitGroup{}
		count int
	)
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		defer wg.Done() // 死锁了
		go func() {
			count += i
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println(demo(1, 2))
}
func demo(a, b int) (c int) {
	defer func() {
		c++ // 这里还是会加加 临时变量
	}()
	c++
	return a + b
}
func testInterface() {
	// s1Val := S1{}
	// s1Ptr := &S1{}
	// s2Val := S2{}
	// s2Ptr := &S2{}

	// var i F
	// i = s1Val
	// i = s1Ptr
	// i = s2Val
	// i = s2Ptr

	// _ = i
}

func testSlice() {
	a := make([]int, 0)
	b := make([]int, 0, 1)
	// cap len %p
	fmt.Println("a.cap:", cap(a), "a.len:", len(a))
	fmt.Println("b.cap:", cap(b), "b.len:", len(b))
	fmt.Printf("a.p = %p b.p= %p\n", a, b)
	a = append(a, 1)
	b = append(b, 1)
	// cap len %p
	fmt.Println("a.cap:", cap(a), "a.len:", len(a))
	fmt.Println("b.cap:", cap(b), "b.len:", len(b))
	fmt.Printf("a.p = %p b.p= %p\n", a, b)
	// len(a)= 1 cap(a)=1
	// len(b)= 1 cap(b)=1
	// a的地址变化 b的地址不变
	a = append(a, 1)
	// cap len
	fmt.Println("a.cap:", cap(a), "a.len:", len(a))
	fmt.Println("b.cap:", cap(b), "b.len:", len(b))
	fmt.Printf("a.p = %p b.p= %p\n", a, b)
	// len(a)= 2 cap(a)=2
	// 求各阶段的值
	// 总结扩容是*2 不是*2+1 cap*2
	// append扩容 地址也会改变
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
