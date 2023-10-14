package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个字符串作为数据源
	data := "Hello, World!"
	
	//实际上就是定义输入源 可以是os.Stdin


	// 使用strings.NewReader创建一个io.Reader实例
	reader := strings.NewReader(data)

	// 创建一个缓冲区，用于存储读取的数据
	buf := make([]byte, 6)
	// 每次只能读取6个
	// 从io.Reader中读取数据并输出
	for {
		n, err := reader.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
