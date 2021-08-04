package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格
func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)            // a b c
	fmt.Println("输入的内容是:", s) // a
}

func useBuffio() {
	var s string
	fmt.Print("请输入内容：")
	reader := bufio.NewReader(os.Stdin) // NewReader参数就是接口类型
	s, _ = reader.ReadString('\n')
	fmt.Printf("输入内容是：%s", s)
}

func main() {
	//useScan()
	useBuffio()
}
