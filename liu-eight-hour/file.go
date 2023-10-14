package main

import "fmt"

var globalVar = 100 // 包级别的变量

var printFunc = func() {
	
	fmt.Println(globalVar)  // 访问包级变量
	fmt.Println("Inside the anonymous function defined outside main")
}

func main() {
    localVar := 200 // 局部变量

    // 匿名函数可以访问和使用 main 函数内的局部变量和包级变量
    func() {
        fmt.Println(localVar)   // 访问局部变量
        fmt.Println(globalVar)  // 访问包级变量
        fmt.Println("Inside the anonymous function")
    }()

    // 匿名函数可以在 main 函数外部定义，并在 main 函数内部使用
    

    printFunc() // 调用匿名函数

    // 修改局部变量的值
    localVar = 300
    printFunc() // 调用匿名函数，会反映新的局部变量值
}
