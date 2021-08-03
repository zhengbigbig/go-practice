package main

import "fmt"

// 使用值接收者和指针接收者区别
type animal interface {
	move()
	eat(s string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口所有方法
//func (c cat) move()  {
//	fmt.Println("miao")
//}
//func (c cat) eat(food string)  {
//	fmt.Println(food)
//}

// 使用指针接收者实现
func (c *cat) move()  {
	fmt.Println("miao")
}
func (c *cat) eat(food string)  {
	fmt.Println(food)
}

func main() {
	var a1 animal
	c1 := cat{"tony",4}
	c2 := &cat{"tom",4}
	a1 = &c1 // 实现animal这个接口的是cat的指针类型，
	a1 = c2
	fmt.Println(a1)
}
