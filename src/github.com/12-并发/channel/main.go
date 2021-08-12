package main

import (
	"fmt"
	"sync"
)

var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan []int // 声明一个传递int切片的通道
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(ch1)     // nil
	ch1 = make(chan int) // 通道初始化,不带缓冲区
	fmt.Println(ch1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch1 // 接收并赋值
		fmt.Println("后台goroutine从通道ch1中取到了", x)
	}()
	ch1 <- 10 // 发送 发送不接收将会hang住
	fmt.Println("10发送到通道ch1")
	fmt.Println(ch1)
	wg.Wait()
}

func bufChannel() {
	fmt.Println(ch1)        // nil
	ch1 = make(chan int, 1) // 带缓冲区的初始化
	fmt.Println(ch1)
	ch1 <- 10
	fmt.Println("10发送到通道ch1了")
	//ch1 <- 20  // 通道缓冲区满了后，再发送将会hang住，导致死锁
	//fmt.Println("20发送到通道cha1了")
	x := <-ch1
	fmt.Println("从通道取出x ", x)
	close(ch1)
}

func main() {
	noBufChannel() // 无缓冲区
	//bufChannel() // 有缓冲区
}
