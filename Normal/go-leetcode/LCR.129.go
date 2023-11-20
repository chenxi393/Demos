package main

// // 写了50分钟 太不熟练了
// // 回溯 + 可行性剪枝
// // 答案其实和我的思路是一样的（要简洁的多）
// // 只是在匹配上将二维字符数组置空 就不用visited数组了

// var gg [][]byte
// var visited [][]bool
// var tt []byte

// func wordPuzzle(grid [][]byte, target string) bool {
// 	gg = grid
// 	tt = []byte(target)
// 	visited = make([][]bool, len(grid))
// 	for i := 0; i < len(grid); i++ {
// 		visited[i] = make([]bool, len(grid[0]))
// 	}
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {
// 			if grid[i][j] == tt[0] {
// 				if findRec(i, j, 1) {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// func findRec(i, j int, fin int) bool {
// 	// 这是成功条件 唯一
// 	if fin == len(tt) {
// 		return true
// 	}
// 	visited[i][j] = true
// 	if j+1 < len(gg[0]) && !visited[i][j+1] && gg[i][j+1] == tt[fin] {
// 		if findRec(i, j+1, fin+1) {
// 			return true
// 		}
// 	}
// 	if i+1 < len(gg) && !visited[i+1][j] && gg[i+1][j] == tt[fin] {
// 		if findRec(i+1, j, fin+1) {
// 			return true
// 		}
// 	}
// 	if j-1 >= 0 && !visited[i][j-1] && gg[i][j-1] == tt[fin] {
// 		if findRec(i, j-1, fin+1) {
// 			return true
// 		}
// 	}
// 	if i-1 >= 0 && !visited[i-1][j] && gg[i-1][j] == tt[fin] {
// 		if findRec(i-1, j, fin+1) {
// 			return true
// 		}
// 	}
// 	visited[i][j] = false
// 	return false
// }
