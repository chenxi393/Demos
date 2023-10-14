package main

import (
    "fmt"
    _ "initLib1" //导入但不适应 使用匿名_  只调用init函数  _在go中表示丢弃 接受但不使用
    "initLib2"
)

func init() {
    fmt.Println("libmain init")
}

func main() {
    fmt.Println("libmian main")

    //initLib1.Test_ini1()
    initLib2.Test_ini2()
}