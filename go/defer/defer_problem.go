package main

import (
	"fmt"
)

func main() {
	fmt.Println("a return:", a())
	fmt.Println("b return:", b())
	fmt.Println("b return:", *c())
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer:", i)
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer:", i)
	}()
	return
}

func c() *int {
	i := new(int)
	defer func() {
		*i++
		fmt.Println("defer:", *i)
	}()
	defer func() {
		*i++
		fmt.Println("defer:", *i)
	}()
	return i
}
