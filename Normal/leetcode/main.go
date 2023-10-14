package main

import "fmt"

func main() {
	//grid := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//target := "ABCCED"
	var root TreeNode
	var ll TreeNode
	var rr TreeNode
	root.Left = &ll
	root.Right = &rr
	root.Val = 1
	ll.Val=2
	rr.Val=3
	fmt.Println(pathTarget(&root,3))
}
