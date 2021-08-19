package main

import (
	"fmt"
	"sort"
)

func main() {
	// 声明切片类型
	var a []string              // 声明一个字符串切片
	var b = []int{}             // 声明一个整型切片并初始化
	var c = []bool{false, true} // 声明一个布尔切片并初始化
	var d = []bool{false, true} // 声明一个布尔切片并初始化
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false
	// 切片是引用类型，不支持直接比较，只能和nil比较

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9}
	a2 := a1[0:4] // 左闭右开
	fmt.Println(a2)
	a3 := a1[:4]
	a4 := a1[3:]
	a5 := a1[:]
	fmt.Println(a3, a4, a5)
	// 切片若从最开始开始，切片容量为数组长度
	// 切片若从中间开始，切片容量为开始到结束的长度
	fmt.Println(len(a3), cap(a3)) // 4 , 5
	fmt.Println(len(a4), cap(a4)) // 2 , 2
	// 切片再切割
	a6 := a5[3:]
	fmt.Println(len(a6), cap(a6))
	// 切片是引用类型，都指向了底层的一个数组
	fmt.Println(a5)
	fmt.Println(a6)
	a5[3] = 100
	fmt.Println(a5)
	fmt.Println(a6)

	// make函数创造切片
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10
	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	// s2=[] len(s2)=0 cap(s2)=10

	// append
	s3 := []string{"a", "b", "c"}
	// s3[3] = "d" // 错误写法，会导致编译错误：索引越界
	// 调用append函数必须用原来的切片变量接收返回值
	s3 = append(s3, "d")
	s4 := append(s3, "d")
	fmt.Println("s3=%v len(s3)=%d cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Println("s4=%v len(s4)=%d cap(s4)=%d\n", s4, len(s4), cap(s4))
	//  [a b c d] 4 6

	// copy
	// 切片是引用类型，实际赋值后两个变量依旧指向同一块内存地址
	// copy()函数可以将一个切片的数据复制到另一个切片空间中
	d1 := []int{1, 2, 3, 4, 5, 6}
	d2 := d1
	d3 := make([]int, 3, 3)
	// func copy(dst, src []Type) int
	// dst 目标切片 src 数据来源切片
	copy(d3, d1)
	fmt.Println(d1, d2, d3)
	// [1 2 3 4 5 6] [1 2 3 4 5 6] [1 2 3]
	d1[0] = 1000
	fmt.Println(d1, d2, d3)
	// [1000 2 3 4 5 6] [1000 2 3 4 5 6] [1 2 3]

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

	var z = make([]int,5,10)
	fmt.Println(z)
	for i := 0; i < 10; i++ {
		z = append(z,i)
	}
	fmt.Println(z)

	var w = [...]int{3,5,6,1,2}
	sort.Ints(w[:])
	fmt.Println(w)
}
