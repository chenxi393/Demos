package main

import (
	"context"
	"fmt"
	"time"
)

func go1(ctx context.Context) {
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("go1: context cancelled")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(i, ctx.Value("key"))
		}

	}
}

// Background 和 TODO 并没有什么不同，只不过用不同的名字来区分不同的场景罢了
func main() {
	ctx := context.Background()

	// context.TODO: 不知道用什么 context 以及不知道需不需要用 context 的时候用
	ctx = context.TODO()
	// 给一个函数方法传递 context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO()；

	//context 是线程安全的，可以放心的在多个 Goroutine 中传递。

	// context.WithValue 传值

	// context.WithCancel 可取消
	ctx, cancel := context.WithCancel(context.WithValue(ctx, "key", "WithCancel"))
	go go1(ctx)
	defer cancel()

	// context.WithDeadline 到指定时间点自动取消（或在这之前手动取消）
	ctx, cancel = context.WithDeadline(context.WithValue(context.TODO(), "key", "WithDeadline"), time.Now().Add(2*time.Second))
	go go1(ctx)

	// context.WithTimeout 一段时间后自动取消（或在这之前手动取消）
	ctx, cancel = context.WithTimeout(context.WithValue(context.TODO(), "key", "WithTimeout"), 3*time.Second)
	go go1(ctx)
	
	time.Sleep(100 * time.Second)
}
