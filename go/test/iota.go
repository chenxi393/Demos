package test_

import (
	"fmt"
)

const (
	_ = iota
	a = 1 << iota
	b
)
const (
	c = iota
	d
)

func Iota() {
	fmt.Println(a, b, c, d)
}
