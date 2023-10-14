package main

import (
	"fmt"
	"reflect"
)

func main(){
	var slice1 []int=make([]int, 1,10)//声明切片方式   数组一般用的少
	slice2 :=[]int{1,2,3,4,5,6}
	slice2 = append(slice2, slice2...)//...表示把slice2作为值展开

	slice1 = append(slice1, slice2...)//扩容会扩容前面的容量的两倍 10变20
	fmt.Printf("cap=%d len=%d val=%v\n",cap(slice1),len(slice1),slice1)
	fmt.Println("------------")
	for _,v :=range slice2{
		fmt.Println(v);
	}
	fmt.Printf("cap=%d len=%d val=%v\n",cap(slice2),len(slice2),slice2)
	/*
	当切片的长度小于 1024 时，每次扩容后的容量将会翻倍。
	当切片的长度大于等于 1024 时，每次扩容后的容量将会增加原容量的 25%，但不会超过所需容量的两倍。
	这似乎是1.8之前的规则 之后是从2倍逐步到1.25倍扩容
	*/


	fmt.Println("------------")
	slice3:=[]int{1,2,3,4,5}
	temp:=slice3[1:3]
	fmt.Println(slice3)
	fmt.Println(temp)
	temp[0]=1
	slice3[1]=111
	fmt.Println(slice3)
	fmt.Println(temp)
	
	
	fmt.Println("------------")
	Mymap:=make(map[string]string)
	Mymap["one"]="java"
	Mymap["two"]="php"
	Mymap["three"]="cpp"
	Mymap["four"]="go"
	fmt.Println(reflect.TypeOf(Mymap),len(Mymap),Mymap)

	
}