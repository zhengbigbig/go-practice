package main

import "fmt"

func main() {
	// 1. &:取地址
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	// 2. *: 根据地址取值
	m := *p
	fmt.Printf("%T\n %v \n", m, m)
	l := &m
	fmt.Printf("%T\n %v \n", l, l)
	// 空指针异常
	//var a *int
	//*a = 100
	//fmt.Println(*a) // nil pointer 内存地址未开辟

	// new函数申请一个内存地址
	var a = new(int)
	fmt.Printf("a: %x \n *a:%d \n", a, *a)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int,10)
	b["hello"] = 100
	fmt.Println(b)
}
