package main

import "fmt"
// 昆仑万维一面
// 0 2 3 0 3
func reverseZero(p []int) []int {
	j := len(p) - 1
	for i := 0; i < len(p) && i < j; i++ {
		if p[i] == 0 {
			for p[j] == 0 && i < j {
				j--
			}
			p[i], p[j] = p[j], p[i]
			j--
		}
	}
	p1 := make([]int, 0, 1)
	mp := make(map[int]struct{}, 2)
	ch := make(chan int, 9)
	pp := new(int)
	ppp := new(map[int]struct{})

	return p
}

func main() {
	p := []int{0, 2, 3, 0, 3}
	fmt.Println(reverseZero(p))
}
