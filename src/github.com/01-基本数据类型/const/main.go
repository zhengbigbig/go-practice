package main

import "fmt"

//const pi = 3.1415
//const e = 2.7182
// 声明这两个常量后，在整个程序运行期间它们的值都不能再发生变化了
// 多个常量一起声明
const (
	pi = 3.1415
	e  = 2.7182
)

// const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
	a4        // 3
)

// 使用_跳过某些值
const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

// iota声明中间插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        //3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1 1 d2 2
	d3, d4 = iota + 1, iota + 2 // d3 2 d3 3
)
const c5 = iota // 0

// 定义数量级
const (
	_  = iota // 0
	KB = 1 << (10 * iota) // 左移10位 2^10 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(b1, b2, b3)
}
