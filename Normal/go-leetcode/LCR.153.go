package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 要注意的点是 会改变底层数据
// 使用切片的指针也可以
func pathTarget(root *TreeNode, target int) [][]int {
	var tt []int
	var ans [][]int
	return ddfs(root, 0, target, tt,ans)
}

func ddfs(r *TreeNode, sum, target int, tt []int, ans [][]int) [][]int {
	if r == nil {
		return ans
	}
	tt = append(tt, r.Val)
	sum += r.Val
	if r.Left == nil && r.Right == nil {
		if sum == target {
			temp := make([]int, 0)     // 这里要注意
			temp = append(temp, tt...) // 不然ans里会共享temp的底层数组
			ans = append(ans, temp)
		}
		return ans
	}
	ans = ddfs(r.Left, sum, target, tt, ans)
	ans = ddfs(r.Right, sum, target, tt, ans)
	return ans
}
