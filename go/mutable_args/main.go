package main

import "fmt"

// 使用 ... （三个点）就可以实现可变参数

// option 之类的可变参数传递 稍有复制 可以看看
func funcName(args ...int) {
	for v := range args {
		fmt.Println(v)
	}
}

func main() {
	funcName(2, 4, 1, 21, 3, 1, 1)

	var s string = "✨😊😂🤣❤️💕😒👍🤦‍♀️🤦‍♂️😎😉😁"
	var ss rune = '💕'  //😂🤣❤️💕😒👍🤦‍♀️🤦‍♂️😎😉😁"
	var sss rune = '😂' //🤣❤️💕😒👍🤦‍♀️🤦‍♂️😎😉😁"

	fmt.Println(s)
	fmt.Printf("%c", ss)
	fmt.Printf("%c", sss)
}
