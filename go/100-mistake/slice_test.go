package main

import "fmt"

func SliceTest() {
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	s1[1] = 1
	print(s2)

	s2 = append(s2, 2)
	print(s2)
	s1 = append(s1, 3) // append的3会覆盖第7行
	print(s1)
	print(s2)

	listing2()
}

func print(s []int) {
	fmt.Printf("len=%d, cap=%d: %v\n", len(s), cap(s), s)
}

func f(s1 []int) {
	s1 = append(s1, 10)
	fmt.Println(s1)
}

func listing2() {
	s := []int{1, 2, 3}
	sCopy := make([]int, 2, 3)
	copy(sCopy, s) //minimum of len(src) and len(dst)

	fmt.Println(sCopy)
	f(sCopy)
	result := append(sCopy, s[2])
	fmt.Println(result)
}
