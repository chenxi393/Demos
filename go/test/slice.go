package test_

import (
	"fmt"
)

func Slice() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func SliceAppend() {
	var nums = make([]int, 0, 4)
	for i := 1; i <= 4; i++ {
		Append(nums, i)
	}
	fmt.Println(nums)
}

func Append(nums []int, num int) {
	nums = append(nums, num)
	// 底层数值也被填充了值，但是对于原nums，len依旧为0，切片可以理解成底层数组的len视图
}
