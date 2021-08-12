package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	x, ok := <-ch1
	fmt.Println(x, ok) // 10 true
	x, ok = <-ch1
	fmt.Println(x, ok) // 20 true
	x, ok = <-ch1
	fmt.Println(x, ok) // 0 false
}
