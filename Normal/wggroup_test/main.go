package main

import (
	"fmt"
	"sync"
)

// 这是牛客的一个题 
// https://www.nowcoder.com/profile/695491024/wrongset/636766954?tags=809
// 测试发现 ans1 ans2 的结果都不是确定的
func main() {
	var wg sync.WaitGroup
	intSlice := []int{}
	for i := 1; i <= 20; i++ {
		intSlice = append(intSlice, i)
	}
	wg.Add(len(intSlice))
	ans1, ans2 := 0, 0
	for _, v := range intSlice {
		vv := v
		go func() {
			defer wg.Done()
			ans1 += v
			ans2 += vv
		}()
	}
	wg.Wait()
	fmt.Printf("ans1:%v,ans2:%v\n", ans1, ans2)
	return
}
