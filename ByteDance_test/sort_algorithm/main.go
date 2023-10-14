package main

import (
	"fmt"
	//"math/rand"
	"sort"
	"time"
)

// 生成有序序列
func generateSortedSeq(size int) []int {
	seq := make([]int, size)
	for i := 0; i < size; i++ {
		seq[i] = i
	}
	return seq
}

// 生成逆序序列
func generateReverseSeq(size int) []int {
	seq := make([]int, size)
	for i := 0; i < size; i++ {
		seq[i] = size - i
	}
	return seq
}


func main() {
	// source := rand.NewSource(time.Now().UnixNano())
	// random := rand.New(source)

	// 生成随机序列
	const size = 10000
	seq := generateSortedSeq(size)
	// for i := 0; i < size; i++ {
	// 	seq[i] = random.Intn(size)
	// }

	// 复制序列，用于三种排序算法的测试
	seq1 := make([]int, size)
	seq2 := make([]int, size)
	benchmarkSeq := make([]int, size)
	copy(seq1, seq)
	copy(seq2, seq)
	copy(benchmarkSeq, seq)

	// 快速排序
	start := time.Now()
	quickSort(seq1)
	elapsed := time.Since(start)
	fmt.Printf("快速排序1完成，耗时：%s\n", elapsed)

	// 插入排序
	start = time.Now()
	quickSort2(seq2)
	elapsed = time.Since(start)
	fmt.Printf("快速排序2完成，耗时：%s\n", elapsed)

	// 基准测试
	fmt.Println("\n基准测试：")
	fmt.Println("排序前：", benchmarkSeq[:10], "...")
	start = time.Now()
	sort.Ints(benchmarkSeq)
	elapsed = time.Since(start)
	fmt.Println("排序后：", benchmarkSeq[:10], "...")
	fmt.Printf("标准库排序完成，耗时：%s\n", elapsed)

	for i := 0; i < size; i++ {
		if !(seq1[i] == seq2[i] && seq2[i] == benchmarkSeq[i]) {
			fmt.Println("排序算法错误")
			break
		}
	}
}
