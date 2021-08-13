package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)  // 必须使用sync.Map内置的Store方法设置键值对
			//value, _ := m.Load(key) // 必须使用sync.Map内置的Load方法设置键值对
			//fmt.Printf("k=%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}
