package main

import "fmt"

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 定义一个bytedancer生成器，输入名称，返回新的用户数据
func genBytedancer(name string) func() (string, int) {
    // 定义字节范分数
    style := 100
    // 返回闭包
    return func() (string, int) {
        // 引用了外部的 style 变量, 形成了闭包
        return name, style
    }
}

// 通过闭包的记忆效应来实现设计模式中工厂模式的生成器
func Factory() {
    // 创建一个bytedancer生成器
    generator := genBytedancer("bytedance001")

    // 返回新创建bytedancer的姓名, 字节范分数
    name, style := generator()
    fmt.Println(name, style)
}


func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	counter := increment()
	// increment()是counter的父函数
	// 因此increment不会被GC 内的局部变量仍然有效
	fmt.Println(counter()) // 输出: 1
	fmt.Println(counter()) // 输出: 2
	fmt.Println(counter()) // 输出: 3
}
