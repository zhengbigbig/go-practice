package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now() // 获取当前时间
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix() // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n",timestamp1)
	fmt.Printf("current timestamp2:%v\n",timestamp2)
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp,0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		timeObj.Year(),timeObj.Month(),timeObj.Day(),timeObj.Hour(),timeObj.Minute(),timeObj.Second())
	fmt.Println(date)
}

func main() {
	timeDemo()
	timestampDemo()
	timestampDemo2(time.Now().Unix())
}
