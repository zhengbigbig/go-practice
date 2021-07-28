package main

import "fmt"

// for 循环标准写法
func forDemo01() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 省略初始语句的for循环写法
func forDemo02() {
	i := 0
	for ; i < 10; i++ { // 简略写法
		fmt.Println(i)
	}
}

// 省略初始语句和结束语句的for循环写法
func forDemo03() {
	i := 0
	for i < 10 { // 相当于其他语言的while语法
		fmt.Println(i)
		i++ // 写法相当于 i = i + 1
	}
}

// 无限循环-死循环
func forDemo04() {
	for {
		// 什么都不写就是最简单的死循环
		// 实际测试这个死循环只会跑满单个CPU内核，推测是针对多核处理器跑死循环任务做了优化
	}
}

func main() {
	forDemo01()
	forDemo02()
	forDemo03()
}
