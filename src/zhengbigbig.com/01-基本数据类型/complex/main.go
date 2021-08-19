package main

import "fmt"

// 复数有实部和虚部，complex64的实部和虚部32位，complex128的实部和虚部为64位
func main() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1,c2)
}
