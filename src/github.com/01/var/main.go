package main

import "fmt"

// 声明变量
//var name string
//var age int
//var isOk bool
// Go语言中推荐使用驼峰式命名
// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "lixiang"
	age = 16
	isOk = true
	// Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk) // 在终端中输入要打印的内容
	fmt.Printf("name:%s", name) // %s 占位符，使用name变量的值去替换
	fmt.Println(age) // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "yyds"
	fmt.Println(s1)
	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面使用
	s3 := "hahaha"
	fmt.Println(s3)
}
