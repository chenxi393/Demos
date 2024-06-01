package fs

// /*
// 这是面试写出来的
// 感觉有点一言难尽
// */

// type FileNode struct {
// 	path     string
// 	children map[string]*FileNode
// 	content  string
// }

// type FileSystem struct {
// 	root *FileNode
// }

// func (f *FileSystem) InitFileSystem() {
// 	root := &FileNode{
// 		path:     "/",
// 		children: make(map[string]*FileNode, 0),
// 	}
// }

// func (f *FileSystem) Write(path, content string) {
// 	// /root/myfile.txt
// 	stringPath := path.split("/")
// 	// _ root myfile.txt
// 	curNode := f.root
// 	for curPath := range stringPath {
// 		if tempNode, ok := curNode[curPath]; !ok {
// 			tempNode.children[curPath] = &FileNode{}
// 		}
// 		curNode = tempNode.children[curPath]
// 	}
// }

// func (f *FileSystem) Read(path string) {

// }

// func (f *FileSystem) Show(path string) {

// }
