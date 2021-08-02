package main

import "fmt"

// 递归：函数自己调用自己
// 递归适合处理那种问题相同，规模越来越小的场景
// 递归一定要有一个明确的退出条件

func f(i int) int {
	if i <= 1 {
		return i
	}
	return i * f(i-1)
}

func f2(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f2(n-1) + f2(n-2)
}

func main() {
	fmt.Println(f(4))
}
