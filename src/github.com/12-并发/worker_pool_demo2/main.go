package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job    *job
	sum int64
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 10)
var resultChan = make(chan *result, 10)

func generate(js chan<- *job) {
	// 循环生成int64类型的随机数，发送到jobChan
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{value: x}
		js <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func sumResult(js <-chan *job, rc chan<- *result) {
	// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	defer wg.Done()
	for {
		job := <-js
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{sum: sum, job: job}
		rc <- newResult
	}
}

func main() {
	wg.Add(1)
	go generate(jobChan)
	wg.Add(24)
	// 开启24个goroutine执行
	for i := 0; i < 24; i++ {
		go sumResult(jobChan, resultChan)
	}
	// 主goroutine从resultChan取出结果并打印到终端输出
	for v := range resultChan {
		fmt.Printf("value:%d sum:%d\n", v.job, v.sum)
	}
	wg.Wait()

}
