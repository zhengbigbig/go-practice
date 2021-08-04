package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("no such file", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err, dstName)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}

func main() {
	_, err := CopyFile("./dst.txt", "./src.txt")
	if err != nil {
		fmt.Println("copy file failed,err:", err)
		return
	}
	fmt.Println("copy done!")
}
