package main

import "fmt"

func ifDemo01(score int) {
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo02(score int) {
	if score := score + 1; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
// 以上两种写法的区别，变量的作用域问题 ！！！
// if 语句外面的变量定义可以用于多个 if 语句
// if 语句里面的变量定义只能用于当前的 if 语句
func main() {
	ifDemo01(50)
	ifDemo02(90)
}
