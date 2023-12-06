package fs

import (
	"fmt"
	"strings"
)

/*
题目 实现一个内存型的文件系统
1. 支持写入内容到对应路径
2. 支持读入指定路径的文件内容
3. 支持展示指定路径下的文件夹和文件名
*/

type fileNode struct {
	path     string
	children map[string]*fileNode
	content  string
}

type FileSystem struct {
	root *fileNode
}

func (f *FileSystem) InitFileSystem() {
	f.root = &fileNode{
		// 这里将默认根目录名设为root
		path:     "root",
		children: make(map[string]*fileNode, 0),
	}
}

func (f *FileSystem) Write(path, content string) {
	// root/myfile.txt
	stringPath := strings.Split(path, "/")
	stringPath = stringPath[1:]
	// root myfile.txt
	curNode := f.root
	for _, curPath := range stringPath {
		_, ok := curNode.children[curPath]
		// 当前路径不存在 逐级创建路径 类似于 mkdir -p
		if !ok {
			curNode.children[curPath] = &fileNode{
				path:     curPath,
				children: make(map[string]*fileNode),
			}
		}
		curNode = curNode.children[curPath]
	}
	curNode.content = content
	fmt.Printf("%s 创建成功\n", path)
}

func (f *FileSystem) Read(path string) {
	stringPath := strings.Split(path, "/")
	stringPath = stringPath[1:]
	curNode := f.root
	for _, curPath := range stringPath {
		tempNode, ok := curNode.children[curPath]
		if !ok {
			// 不存在返回错误信息
			fmt.Printf("%s 不存在\n", path)
			return
		}
		curNode = tempNode
	}
	fmt.Printf("%s\n", curNode.content)
}

func (f *FileSystem) Show(path string) {
	stringPath := strings.Split(path, "/")
	stringPath = stringPath[1:]
	curNode := f.root
	for _, curPath := range stringPath {
		tempNode, ok := curNode.children[curPath]
		if !ok {
			// 不存在返回错误信息
			fmt.Printf("%s 不存在\n", path)
			return
		}
		curNode = tempNode
	}
	for _, ch := range curNode.children {
		fmt.Printf("%s ", ch.path)
	}
	fmt.Println()
}
