package main

// 和题解类似的代码 非常简洁
func wordPuzzle(grid [][]byte, target string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if findRec(i, j, 0, grid, target) {
				return true
			}
		}
	}
	return false
}

func findRec(i, j int, fin int, gg [][]byte, tt string) bool {
	if i < 0 || j < 0 || j == len(gg[0]) || i == len(gg) || tt[fin] != gg[i][j] {
		return false
	}
	if fin == len(tt)-1 {
		return true
	}
	gg[i][j] = 0
	fin++
	ans := findRec(i, j+1, fin, gg, tt) || findRec(i+1, j, fin, gg, tt) ||
		findRec(i, j-1, fin, gg, tt) || findRec(i-1, j, fin, gg, tt)
	gg[i][j] = tt[fin-1]
	return ans
}
