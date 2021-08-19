package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口嵌套
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string2 string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了mover接口
func (c *cat) move() {
	fmt.Println("🐈")
}

// cat 实现了eater接口
func (c *cat) eat(food string) {
	fmt.Println(food)
}

func main() {
	var a1 animal
	c1 := cat{
		name: "tony",
		feet: 0,
	}
	a1 = &c1
	a1.move()
}
