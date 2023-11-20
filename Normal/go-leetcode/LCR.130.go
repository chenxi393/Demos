package main

// 这题目描述有点傻逼
// 乍一看是只能向右或者下 一条路走到黑
// 然后求整理最大的格子数
// 但是题解是两个方向的和 就是在30行的条件下求可达格子的总数
var a, b int
var vis [][]bool

func wardrobeFinishing(m int, n int, cnt int) int {
	a, b = m, n
	sum := 0
	vis = make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	dfs(0, 0, cnt, &sum)
	return sum
}
func digit(t int) int {
	sum := 0
	for t > 0 {
		sum += t % 10
		t /= 10
	}
	return sum
}

func dfs(i, j, cnt int, sum *int) {
	if i == a || j == b || digit(i)+digit(j) > cnt || vis[i][j] {
		return
	}
	*sum++
	vis[i][j] = true
	dfs(i, j+1, cnt, sum)
	dfs(i+1, j, cnt, sum)
}
