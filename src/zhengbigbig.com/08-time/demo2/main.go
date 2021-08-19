package main

import (
	"fmt"
	"time"
)

func tickDemo() {
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) // 每次都会执行的任务
	}
}

func formatDemo() {
	now := time.Now()
	// 格式化的模块为go的出生时间2006年1月2号15点04分
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Printf("%T %v\n",timeObj,timeObj)
}

func formatTimeZone() {
	// 时区
	now := time.Now()
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05","2019-08-04 14:42:11")
	// 按照东八区的时区和格式解析一个字符串格式的时间
	// 根据字符串加载时区
	loc,err:=time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed,err:",err)
		return
	}
	// 按照指定时区解析时间
	fmt.Println(loc)
	timeObj,err := time.ParseInLocation("2006-01-02 15:04:05","2021-08-06 14:42:11",loc)
	if err != nil {
		fmt.Println("parse time failed,err:",err)
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	//now := time.Now()
	//last := now.Add(-24 * time.Hour)
	//fmt.Println(now.Sub(last).Hours()) // 24
	//fmt.Println(now.Equal(last)) // false
	//fmt.Println(now.Before(last)) // false
	//fmt.Println(now.After(last)) // true
	//tickDemo()
	//formatDemo()
	formatTimeZone()
}
