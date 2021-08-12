package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once
func f1 (ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1,ch2 chan int)  {
	defer wg.Done()
	for {
		i, ok := <-ch1
		if !ok { // 当关闭的时候ok 为false
			break
		}
		ch2 <- i * i
	}
	once.Do(func() {
		close(ch2)
	}) // 确保只关闭一次，不然会导致panic
}

func main() {
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 100)
	wg.Add(3)
	// 开启goroutine将0~100的数发送到ch1中
	go f1(ch1)
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	// 定义一个生产者，两个消费者
	go f2(ch1,ch2)
	go f2(ch1,ch2)
	wg.Wait()
	// 在主goroutine中从ch2中接收值并打印
	for i := range ch2 { // 关闭之后可读不可写
		fmt.Println(i)
	}
}
