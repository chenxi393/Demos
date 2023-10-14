package main

import (
	"log"
	"os"
)

func main() {
	//这里
	// 创建一个log.Logger实例，将日志输出到标准输出
	logger := log.New(os.Stdout, "Example: ", log.LstdFlags)

	// 记录一条普通日志
	logger.Println("This is a normal log message.")

	// 记录一条带有格式化字符串的日志
	name := "Alice"
	age := 30
	logger.Printf("User %s is %d years old.", name, age)

	// 记录一条带有日期和时间戳的日志
	logger.Println("This is another log message.")
}
