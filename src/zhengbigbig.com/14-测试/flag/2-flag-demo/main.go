package main

import (
	"flag"
	"fmt"
)

// flag获取命令行参数
func main() {
	// 创建一个标志位参数
	name := flag.String("name","张三","姓名")
	age := flag.Int("age",18,"年龄")
	married := flag.Bool("married",false,"婚否")
	delay := flag.Duration("d",0,"时间间隔")
	// 使用flag
	flag.Parse()
	fmt.Println(*name,*age,*married,*delay)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

}
