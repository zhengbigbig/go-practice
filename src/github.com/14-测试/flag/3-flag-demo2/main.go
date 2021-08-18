package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	flag.Parse()
	fmt.Println(name, age, married, delay)
}
