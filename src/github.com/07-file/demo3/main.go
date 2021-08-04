package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio 按行读取
func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 注意是字符
		if err == io.EOF {
			fmt.Println("read successful")
			return
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		fmt.Print(line)
	}
}
