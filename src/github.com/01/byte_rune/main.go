package main

import (
	"fmt"
)

// 遍历字符串 traversalString
func main() {
	s := "hello世界"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	// 字符串修改
	s1 := "aaa"
	s2 := []rune(s1) // 把字符串强制转换成一个rune切片
	s2[0] = 'b'
	fmt.Println(string(s2)) // 把rune切片强制转换成字符串

	c1 := "牛"
	c2 := '牛'
	fmt.Printf("c1:%T c2:%T", c1, c2) //c1:string c2:int32
}
