package main

import (
	"fmt"
	"os"
	"syscall"
)

func mmap_write() {
	// mmap + write 使用

	map_file, err := os.Create("./input.txt")
	map_file.WriteString("---------------------")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mmap, err := syscall.Mmap(int(map_file.Fd()), 0, 20, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	// 解释一下参数 文件的描述符 偏移量0 映射文件的大小 权限 flag的参数
	if err != nil {
		fmt.Println("无法将文件映射到内存:", err)
		os.Exit(1)
	}

	content := []byte("Hello, World!")
	copy(mmap, content)
	// 执行数据同步，确保数据持久化到磁盘

	os.WriteFile("./result.txt", mmap, 0666)

}

func main() {
	mmap_write()

	// 代码部分
	// io.Copy(dst, src)
	// copyBuffer -> ReadFrom -> ReaderFrom interface(很多包都实现了这个接口)
	// TCPConn.readFrom 进行分级调用 splice?:sendfile?:genericReadFrom
}
