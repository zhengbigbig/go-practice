package main

import (
	"fmt"
)

func main() {
	// 从切片中删除元素
	e := []int{1,2,4,5,6,7}
	e = append(e[:2],e[3:]...)
	fmt.Println(e)

	x := [...]int{1,3,5}
	y := x[:]
	fmt.Println(y,len(y),cap(y)) // [1 3 5] 3 3
	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存
	y = append(y[:1],y[2:]...)
	fmt.Println(y,len(y),cap(y)) // [1 5] 2 3
	fmt.Println(x) // [1 5 5]

	//var z = make([]int,5,10)
	//fmt.Println(z)
	//for i := 0; i < 10; i++ {
	//	z = append(z,i)
	//}
	//fmt.Println(z)
	//
	//var w = [...]int{3,5,6,1,2}
	//sort.Ints(w[:])
	//fmt.Println(w)

}