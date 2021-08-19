package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. str -> int
	str := "10000"
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("str parse int failed,err:",err)
		return
	}
	fmt.Println(parseInt)

	retInt,_ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n",retInt,retInt)
	// 2. int -> str
	i := int32(2316)
	//fmt.Println(string(i)) // 实际为ascii码，要转为字符串"2316"
	fmt.Println(strconv.Itoa(int(i)))
	fmt.Println(fmt.Sprintf("%d",i))
	// 3. str -> bool
	boolStr := "true"
	parseBool, _ := strconv.ParseBool(boolStr)
	fmt.Println(parseBool)
}
