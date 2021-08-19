package main

import "fmt"

// 变量的作用域
var x = 100 // 定义一个全局变量

// 定义一个函数
func f1() {
	// x := 10
	// 函数中查找变量的顺序
	// 1. 先在函数内部查找
	// 2. 找不到就往函数的外面查找，一直找到全局
	fmt.Println(x)
	if i:= 1;i > 1 {
		fmt.Println(i)
	}
	// 外面无法访问i
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	// 外面无法访问j
}

func main() {
	f1()
}
