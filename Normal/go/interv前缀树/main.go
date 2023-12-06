package main

import (
	"fmt"
	"interview/fs"
)

func main() {
	var fileSystem fs.FileSystem
	fileSystem.InitFileSystem()
	var k int
	for {
		fmt.Printf("1. Insert a file in a path.\n")
		fmt.Printf("2. Read a file in a path.\n")
		fmt.Printf("3. Show files in a path.\n")
		fmt.Scanln(&k)
		switch k {
		case 1:
			{
				fmt.Println("Input a path and file content such as root/test/test.txt 这是一个测试文件")
				var path, content string
				fmt.Scanf("%s %s\n", &path, &content)
				fileSystem.Write(path, content)
			}
		case 2:
			{
				fmt.Println("Input a path such as root/test/test.txt")
				var path string
				fmt.Scanln(&path)
				fileSystem.Read(path)
			}
		case 3:
			{
				fmt.Println("Input a path such as root, root/test")
				var path string
				fmt.Scanln(&path)
				fileSystem.Show(path)
			}
		default:
			fmt.Println("Input error")
		}
	}
}
