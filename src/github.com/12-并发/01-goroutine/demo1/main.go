package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(6) // 默认CPU逻辑核心数，默认跑满整个CPU
	last := time.Now()
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	now := time.Now()
	fmt.Println(now.Sub(last))
}
