package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T", a1, a2)

	// 数组的初始化
	// 如果不初始化，默认元素都是零值，(布尔值：false，整型和浮点型都是0，字符串："")
	// 1.
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.
	//a10 := [9]int{0,1,2,3,4,5,6,7}
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a10)
	// 3. 根据索引来初始化
	a3 := [...]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2. for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	// 多维数组的遍历
	for _, v := range a11 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2) // [1 2 3] [100 2 3]

	b3 := [...]int{1, 3, 5, 7, 8}
	v3 := 0
	for _, v := range b3 {
		v3 += v
	}
	fmt.Println(v3, "v3")

	for i := 0; i < len(b3); i++ {
		for j := i + 1; j < len(b3); j++ {
			if b3[i]+b3[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
