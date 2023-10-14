package main

import (
	"fmt"
	"unsafe"
)

// 定义一个接口
type TestInterface interface {
	TestMethod()
}

// 定义一个结构体类型
type MyStruct struct {
	Name string
}

// 实现 TestInterface 接口的 TestMethod 方法
func (m MyStruct) TestMethod() {
	fmt.Println("TestMethod called")
}

func main() {
	// 示例1：使用 eface
	var e interface{}
	myStruct := MyStruct{Name: "John"}
	e = myStruct

	// 获取 eface 的类型信息结构体指针
	typeInfo := *(**struct {
		_   uintptr
		typ uintptr
	})(unsafe.Pointer(&e))

	// 获取 eface 的值空间指针
	value := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&e)) + unsafe.Sizeof(uintptr(0))))

	fmt.Printf("eface type: %T\n", typeInfo)
	fmt.Printf("eface value: %v\n", *(*MyStruct)(value))

	// 示例2：使用 iface
	var i TestInterface
	i = myStruct

	// 获取 iface 的 itab 结构指针
	itab := *(*struct {
		_    uintptr
		_    uintptr
		typ  uintptr
		data uintptr
	})(unsafe.Pointer(&i))

	// 获取 iface 的值空间指针
	data := *(*unsafe.Pointer)(unsafe.Pointer(itab.data))

	fmt.Printf("iface static type: %T\n", itab.typ)
	fmt.Printf("iface dynamic type: %T\n", *(*MyStruct)(data))
}