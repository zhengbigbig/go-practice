package main

import "fmt"

func main() {
	ch := make(chan int,1)
	// 不是从上到下判断，而是随机
	for i := 0; i < 10; i++ {
		select {
		case x:=<-ch:
			fmt.Println(x,"取")
		case  ch <- i:
			fmt.Println(i,"存")
		}
	}
}
