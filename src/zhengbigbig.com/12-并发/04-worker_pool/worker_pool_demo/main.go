package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	// 5个任务到3个goroutine中执行
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	wg.Add(3)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	close(jobs) // 需要告诉接收值的channel停止等待
	wg.Wait()
	// 输出结果
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
