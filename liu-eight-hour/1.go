package main

import (
	"fmt"
)

func max(s1 int, s2 int) int {
	if s1 >= s2 {
		return s1
	}
	return s2
}

func maxSubArray(nums []int) int {
	maxsum, cursum := nums[0], 0
	for i := range nums {
		cursum = max(cursum+nums[i], nums[i])
		maxsum = max(cursum, maxsum)
	}
	return maxsum
}

func main() {
	a := []int{-2, 2, -1, 4}
	fmt.Println(maxSubArray(a))
}
