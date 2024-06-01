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
			ans1 += v  // v的作用域是整个for循环期间 似乎go 1.22 更改了v的作用域
			ans2 += vv // 这里不是原子相加
			//fmt.Printf("v = %d vv=%d\n", v, vv)
			// 但是这里打印vv 肯定是都有的了 可以试试 因为vv的作用域是单次循环内
		}()
		// 24.2.22 cc 额 这里相加肯定是不
		// 不是原子的啊 不一定加出想要的结果
	}
	wg.Wait()
	fmt.Printf("ans1:%v,ans2:%v\n", ans1, ans2)
}
