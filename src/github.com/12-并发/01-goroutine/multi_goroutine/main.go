package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 这里使用sync.WaitGroup来实现goroutine同步
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func main() {
	//rand.Seed(time.Now().UnixNano()) // 保证每次执行的时候都是不一样的
	for i := 0; i < 10; i++ {
		//r := rand.Intn(10) // 0<=x<10
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有等级的goroutine都结束，计算器减为0
}
