package main

import (
	"fmt"
	"unsafe"
)

const (//iota 只能配合const()使用
	BEIJING = 11 * iota
	SHANGHAI
	SHENZHEN
	DALIAN
	CHANGESHA
)

// const(
// 	a,b=iota+1,iota+2   //iota=0
// 	c,d

// 	e,f=iota*2,iota*3
// 	g,h
// )

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
	//字符串类型在 go 里是个结构,
	//包含指向底层数组的指针和长度,
	//这两部分每部分都是 8 个字节，
	//所以字符串类型大小为 16 个字节。
)

func main() {
	const length = 10

	fmt.Println(length)

	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)
	fmt.Println(DALIAN)
	fmt.Println(CHANGESHA)
	println(a, b, c)
	// fmt.Println(a,b,c,d,e,f,g,h)
}
