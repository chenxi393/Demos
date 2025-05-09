package test_

import (
	"fmt"
)

func RangeV() {
	m := make(map[int]*int)
	s := []int{1, 2, 3}
	for _, v := range s {
		m[v] = &v
	}
	fmt.Println(*m[2])
}
