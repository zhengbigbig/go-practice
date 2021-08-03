package main

import "fmt"

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量，都是speaker类型
}

type cat struct {
}
type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵")
}
func (d dog) speak() {
	fmt.Println("汪")
}
func da(x speaker) {
	// 接收一个参数，传递进来，就打印什么
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)

	var ss speaker // 定义一个接口类型：speaker 变量:ss
	ss = c1
	ss = d1
	fmt.Println(ss)

}
