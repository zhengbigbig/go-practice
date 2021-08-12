package main

import (
	"fmt"
	"time"
)

var notifyCh = make(chan struct{}, 5)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
		notifyCh <- struct{}{} // 使用匿名结构体减少内存占用
	}
}

// 全部异步执行
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启5个任务
	go func() {
		for i := 0; i < 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	// 开启3个goroutine
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}
	// 控制结束
	go func() {
		for i := 0; i < 5; i++ {
			<-notifyCh
		}
		close(results)
	}()
	// 输出结果
	for result := range results { // results关闭后才结束
		fmt.Println(result)
	}

}
