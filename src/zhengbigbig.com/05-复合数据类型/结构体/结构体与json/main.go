package main

import (
	"encoding/json"
	"fmt"
)

// Person 结构体与JSON
// 1. 序列化 Go结构体变量 -> json格式字符串
// 2. 反序列化 json格式字符串 -> Go结构体变量
// json包 大写才能拿到，若前端需要返回小写，可以在结构体中加tag
type Person struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := Person{"zbb", 12}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,error:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	// 反序列化
	str := `{"name":"zbb","age":18}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
