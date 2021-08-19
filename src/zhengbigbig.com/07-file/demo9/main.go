package main

import (
	"fmt"
	"io"
	"os"
)

// 往文件中间插入内容
func f(path string) {
	// 1. 打开要操作的文件
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	// 2. 创建临时文件，直接写会覆盖，因此需要借助一个临时文件，
	//file.Write([]byte{'d'})
	tmpFile,err := os.OpenFile("./test.tmp",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("create tmp file failed,err:",err)
		return
	}
	// 3. 读取文件将光标前写入临时文件
	var ret [1]byte
	n ,err := file.Read(ret[:])
	if err != nil {
		fmt.Println("read file failed,err:",err)
		return
	}
	tmpFile.Write(ret[:n])
	// 4. 光标移动
	file.Seek(1, 0) // 光标移动
	// offset 相对偏移量
	// whence 相对位置 0 相对文件开头 1 相对当前位置 2 相对文件结尾
	// 返回新的偏移量(相对开头)和可能的错误
	//fmt.Println(string(ret[:n]))
	// 5. 写入要插入的内容
	tmpFile.Write([]byte{'x'})
	// 6. 将源文件后续内容也写入临时文件
	var x [1024]byte
	for {
		n ,err := file.Read(x[:])
		if err == io.EOF {
			fmt.Println("read done!")
			break
		}
		if err != nil {
			fmt.Println("read failed,err:",err)
			return
		}
		tmpFile.Write(x[:n])
	}
	// 7. 删除源文件，将临时文件更换为源文件
	file.Close()
	tmpFile.Close()
	os.Rename(tmpFile.Name(),file.Name())

}

func main() {
	f("./test.txt")
}
