package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 10
	var str string = "Hello"
	var flag bool = true
	var array [10]int
	array2:=[7]int{1,2,3,4,5}

	fmt.Println("type is ",reflect.TypeOf(array))
	fmt.Println("type is ",reflect.TypeOf(array2))
	fmt.Printf("num's type: %T\n", num)
	fmt.Println("str's type:", reflect.TypeOf(str))
	fmt.Println("flag's type:", reflect.TypeOf(flag))

	for i:=0;i<10;i++{
		fmt.Println(array[i])
	}

	for _,v := range array2{
		fmt.Println(v);
	}
}
