package main

import (
	"fmt"
	"strings"
)

// 闭包
// 函数内部可以访问函数外部作用域变量
func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
}

// 定义一个函数对f2进行包装
func f3(f func(int, int)) func() {
	tmp := func() {
		fmt.Println("hello")
		f(1, 2)
	}
	return tmp
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("test"))
}
