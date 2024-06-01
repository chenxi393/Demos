package main

import (
    "os"
    "fmt"
    "runtime/trace"
)

func main() {
    // go tool trace trace.out 
    // 记录运行时GMP
    //创建trace文件
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    //启动trace goroutine
    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()

    //main
    fmt.Println("Hello World")
}