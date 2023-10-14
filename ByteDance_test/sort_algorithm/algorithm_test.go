package main

import (
	"math/rand"
	"sort"

	// "sort"
	"testing"
	"time"
)

const size = 5

func BenchmarkSort1(b *testing.B) {
	// 创建独立的随机生成器
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < b.N; i++ {
		// 生成随机数组
		nums := make([]int, size)
		for j := 0; j < size; j++ {
			nums[j] = random.Intn(size)
		}

		// 调用需要进行基准测试的函数
		quickSort(nums)
	}
}

func BenchmarkSort2(b *testing.B) {
	// 创建独立的随机生成器
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < b.N; i++ {
		// 生成随机数组
		nums := make([]int, size)
		for j := 0; j < size; j++ {
			nums[j] = random.Intn(size)
		}

		// 调用需要进行基准测试的函数
		heapSort(nums)
	}
}

func BenchmarkSort3(b *testing.B) {
	// 创建独立的随机生成器
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < b.N; i++ {
		// 生成随机数组
		nums := make([]int, size)
		for j := 0; j < size; j++ {
			nums[j] = random.Intn(size)
		}

		// 调用需要进行基准测试的函数
		sort.Ints(nums)
	}
}

func BenchmarkSort4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateReverseSeq(size)
		quickSort2(nums)
	}
}

func BenchmarkSort6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateSortedSeq(size)
		quickSort2(nums)
	}
}
