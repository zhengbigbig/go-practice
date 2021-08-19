package main

import "fmt"

// 算数运算
func suanshuDemo() {
	n1 := 20
	n2 := 5
	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	n1++
	n2--
	fmt.Println(n1)
	fmt.Println(n2)
}

// 关系运算
func guanxiDemo() {
	n1 := 20
	n2 := 5
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
}

// 逻辑运算
func luojiDemo() {
	n1 := true
	n2 := false
	fmt.Println(n1 && n2)
	fmt.Println(n1 || n2)
	fmt.Println(!n1)
	fmt.Println(!n2)
}

// 位运算 1101--13 11--3
func weiDemo() {
	n1 := 13
	n2 := 3
	fmt.Printf("13的二进制为%b\n", n1)
	fmt.Printf("3的二进制为%b\n", n2)
	fmt.Printf("13和3进行位运算 &,结果为 %b,%v\n", n1&n2, n1&n2)
	fmt.Printf("13和3进行位运算 |,结果为 %b,%v\n", n1|n2, n1|n2)
	fmt.Printf("13和3进行位运算 ^,结果为 %b,%v\n", n1^n2, n1^n2)
	fmt.Printf("3进行位运算 << 左移10位,结果为 %b,%v\n", 3<<10, 3<<10)
	fmt.Printf("3进行位运算 >> 右移1位,结果为 %b,%v\n", 3>>1, 3>>1)
}

// 赋值运算
func fuzhiDemo() {
	num := 10
	x := 2
	num = num * x
	fmt.Println(num)
	num /= x
	fmt.Println(num)
}

func main() {
	// suanshuDemo()
	// guanxiDemo()
	// luojiDemo()
	// weiDemo()
	// fuzhiDemo()
}