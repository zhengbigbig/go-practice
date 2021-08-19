package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2() int {
	return 10
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数还可以作为返回值
func f4(x func() int) func(int, int) int {
	// 匿名函数，一般函数内部无法声明带名字的函数
	// 1. 如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)
	// 2. 返回匿名函数
	return func(a int, b int) int {
		return a + b + x()
	}
}


func main() {

}
