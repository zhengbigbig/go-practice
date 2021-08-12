package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendNum(ch chan int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(time.Second * 5)
	}
}

func main() {
	ch := make(chan int, 1)
	// ch <- 100 // 发送一个值到通道
	// <- ch // 将通道中100取出来
	//x,ok := <-ch // 再取则会阻塞
	go sendNum(ch)
	for {
		x, ok := <-ch // 没有值时，就会阻塞，一直等待
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}
}
