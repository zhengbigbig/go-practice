package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello single_goroutine",i)
}

// 程序启动之后会创建一个goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i) // 开启一个单独的goroutine去执行hello函数 （任务）
		//go func() {println(i)}() // 匿名函数  当前i是由函数外部来的，闭包中的i还未变化，可能for循环已经跑了好几个了，因此结果会出现打印多个相同的值
		go func(i int) {println(i)}(i) // 将i传递进去，拿到的i是参数的i，不是外面的i

	}
	fmt.Println("main single_goroutine done")
	// main函数结束，由main函数启动的goroutine也都结束了，因此需要等待
	time.Sleep(time.Second)
}
