package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读方式打开当前目录下的main.go文件
	// 相对路径，需要编译后才能识别，在goland运行不是此路径
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer fileObj.Close()
	// 使用Read方法读取数据
	var content []byte
	var tmp = make([]byte, 767)
	for {
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			fmt.Println("read successful")
			break
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			break
		}
		fmt.Printf("读了%d个字节\n",n)
		content = append(content,tmp[:n]...)
		if n < 767 {
			break
		}
	}
	fmt.Println(string(content))
}
