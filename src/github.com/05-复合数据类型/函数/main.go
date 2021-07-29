package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}

// 无命名的返回值
func f() int {
	return 3
}

// 有命名的返回值
func f1() (ret int) {
	ret = 3
	return
}

// 返回值可以命名也可不命名
// 使用命名返回值可以省略return后面

// 多个返回值
func f2() (int, string) {
	return 1, "2"
}

// 参数类型简写：当连续多个参数的类型一致时，可以将非最后一个参数类型省略
func f3(x, y, z int, m, n string, i, j bool) int {
	return x + y + z
}

// 可变长参数
func f4(x string, y ...int) (string, []int) {
	fmt.Println(x)
	fmt.Println(y) // y类型是切片 []int
	return x, y
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
	_, n := f2()
	fmt.Println(n)
	x, y := f4("1", 1, 23, 4)
	fmt.Println(x, y)

}
