package main


/*
理清思路再写 不然凭感觉和样例 一步步试很大概览不对
若树 B 是树 A 的子结构，则子结构的根节点可能为树 A 的任意一个节点。
1. 遍历A的每一个结点
2. 对A的每一个结点生成的子树都检查是不是
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	// 实质对A做先序遍历
	if helper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B) {
		return true
	}
	return false
}

func helper(A *TreeNode, B *TreeNode) bool {
	if B == nil && A == nil { //说明匹配完成
		return true
	}
	if A == nil || B == nil ||A.Val != B.Val{ //匹配失败
		return false
	}
	return helper(A.Left, B.Left) && helper(A.Right, B.Right)
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	queue := []*TreeNode{}
	root := &TreeNode{Val: nums[0]}
	queue = append(queue, root)

	i := 1
	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) {
			if nums[i] != -1 {
				node.Left = &TreeNode{Val: nums[i]}
				queue = append(queue, node.Left)
			}
			i++
		}

		if i < len(nums) {
			if nums[i] != -1 {
				node.Right = &TreeNode{Val: nums[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}
