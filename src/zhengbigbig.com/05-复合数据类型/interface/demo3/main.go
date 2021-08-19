package main

import "fmt"

// 空接口
// interface 关键字
// interface{} 空接口类型
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zbb"
	m1["age"] = 100
	m1["sex"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}

	var x interface{}
	x = "hello"
	v, ok := x.(string)
	if !ok {
		fmt.Printf("断言失败，实际类型为%T\n", x)
	} else {
		fmt.Printf("值为：%v\n", v)
	}
	assign(x)
}

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)
	case bool:
		fmt.Println("bool", t)
	}
}
