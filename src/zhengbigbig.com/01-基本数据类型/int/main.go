package main

import "fmt"

func main() {
	// 定义十进制数
	a := 10
	fmt.Printf("%d \n\n", a) // 占位符 %d表示十进制显示为 10
	fmt.Printf("%b \n\n", a) // 占位符 %b表示二进制显示为 1010

	// 定义八进制数，需要以0开头
	b := 077
	fmt.Printf("%o \n\n", b) // 占位符 %o表示八进制显示为 77
	// 定义十六进制数，需要以0x开头
	c := 0xff
	fmt.Printf("%x \n\n", c) // 占位符 %x表示十六进制显示为 ff，小写字母显示
	fmt.Printf("%X \n\n", c) // 占位符 %x表示十六进制显示为 FF，大写字母显示

	// 声明int8类型的变量
	i8 := int8(9) // 明确指定int8类型，否则就默认为int类型
	fmt.Printf("%T \n", i8)
}
