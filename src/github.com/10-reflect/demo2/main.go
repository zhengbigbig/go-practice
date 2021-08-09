package main

import (
	"fmt"
	"reflect"
)

type Kind uint
const (
	Invalid Kind = iota // 非法类型
	Bool				// 布尔型
	Int					// 有符号整型
	Int8 				// 有符号8位整型
	Int16				// 有符号16位整型
	Int32				// 有符号32位整型
	Int64				// 有符号64位整型
	Uint				// 无符号整型
	Uint8				// 无符号8位整型
	Uint16				// 无符号16位整型
	Uint32				// 无符号32位整型
	Uint64				// 无符号64位整型
	Uintptr				// 指针
	Float32				// 单精度浮点数
	Float64				// 双精度浮点数
	Complex64			// 64位复数类型
	Complex128			// 128位复数类型
	Array				// 数组
	Chan				// 通道
	Func				// 函数
	Interface			// 接口
	Map					// 映射
	Ptr					// 指针
	Slice				// 切片
	String				// 字符串
	Struct				// 结构体
	UnsafePointer		// 底层指针
)

type myInt int64
func reflectType(x interface{})  {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n",t.Name(),t.Kind())
}

func main() {
	var a *float32 // 指针
	var b myInt // 自定义类型
	var c rune // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type: myInt kind:int64
	reflectType(c) // type: int32 kind: int32

	type person struct {
		name string
	}
	reflectType(person{name: "xxx"}) // type:person kind:struct
}
