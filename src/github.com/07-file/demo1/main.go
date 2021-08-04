package main

import (
	"fmt"
	"os"
)

func main() {
	// 只读方式打开当前目录下的main.go文件
	// 相对路径，需要编译后才能识别，在goland运行不是此路径
	file,err := os.Open("./main.go")
	if err != nil{
		fmt.Println("open file failed,error",err)
		return
	}
	fmt.Println(file.Name())
	// 关闭文件
	//file.Close()
	defer file.Close()
}
