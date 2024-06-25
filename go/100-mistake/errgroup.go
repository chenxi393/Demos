package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func handler2(ctx context.Context, strs []string) ([]string, error) {
	results := make([]string, len(strs))
	g, ctx := errgroup.WithContext(ctx) // 可以取消上下文
	//g := errgroup.Group{} // 无法取消
	for i, circle := range strs {
		i := i
		str := circle
		g.Go(func() error {
			result, err := foo(ctx, str, i)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("err:", err)
		return nil, fmt.Errorf("handler return err, %w", err)
	}
	return results, nil
}

func foo(ctx context.Context, str string, i int) (string, error) {
	if i == 2 {
		return "", fmt.Errorf("i:%d, err", i)
	}
	time.Sleep(time.Second * 2)

	select {
	case <-ctx.Done(): // 可以监听到取消 从而退出
		return "", nil
	default:
		fmt.Println("no error")
		return str, nil
	}
}
func Test_errgroup() {
	strings, err := handler2(context.TODO(), []string{"1", "2", "3"})
	if err != nil {
		fmt.Println("", err.Error())
		return
	}
	fmt.Println(strings)
}
