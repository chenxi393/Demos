package main

import (
	"bufio"

	"math/rand"
	"os"
	"strings"

	"github.com/bytedance/gopkg/lang/fastrand"
)

// 1. unit_test
func HellTom() string {
	return "Tom"
}

// 2. mock_test
func ReadFirstLine() string {
	open, err := os.Open("log")

	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	//NewScanner 默认读取一行(defaults to ScanLines.)
	for scanner.Scan() { // 从中读取一行
		return scanner.Text() //获取读取的文本
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destline := strings.ReplaceAll(line, "11", "00")
	return destline
}

// 3. benchmark_test
var ServerIndex [10]int

func InitSeverIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int {
	return ServerIndex[rand.Intn(10)]
}

func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}

// 4.test for 性能
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// 5. slice预分配
func NoPreAlloc(size int) {
	data := make([]int, 0)
	for k := 0; k < size; k++ {
		data = append(data, k)
	}
}

func PreAlloc(size int) {
	data := make([]int, 0,size)
	for k := 0; k < size; k++ {
		data = append(data, k)
	}
}