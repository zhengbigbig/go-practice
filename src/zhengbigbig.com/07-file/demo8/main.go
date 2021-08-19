package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func fileRead(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	content := make([]byte, 10)
	n, err := file.Read(content)
	if err == io.EOF {
		fmt.Println("file read done", err)
		return
	}
	if err != nil {
		fmt.Println("read failed,err:", err)
		return
	}
	fmt.Printf("读取内容数量：%d\n", n)
	fmt.Println(string(content))
}

func fileReadCir(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	content := make([]byte, 100)
	var ret []byte
	for {
		n, err := file.Read(content)
		if err == io.EOF {
			fmt.Println("read done!", err)
			break
		}
		if err != nil {
			fmt.Println("read failed,err:", err)
			return
		}
		ret = append(ret, content[:n]...)
	}
	fmt.Println(string(ret))
}

func fileReadBufio(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var ret []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("read done!", err)
			break
		}
		if err != nil {
			fmt.Println("read failed,err:", err)
			return
		}
		ret = append(ret, string(line))
	}
	fmt.Println(strings.Join(ret, "\n"))
}

func fileReadIoUtil(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Println(string(content))
}

func fileWrite(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	str := "fuck abc"
	file.Write([]byte(str))
	file.WriteString("stay foolish")
}

func fileWriteBufio(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		writer.WriteString(str)
	}
	writer.Flush()
}

func fileWriteIoUtil(path string) {
	str := "hello"
	err := ioutil.WriteFile(path, []byte(str), 0644)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}

func copyFile(srcName, dstName string) (written int64, err error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open srcFile file failed,err:", err)
		return
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open dstFile file failed,err:", err)
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func main() {
	//fileRead("./test.txt")
	//fileReadCir("./test.txt")
	//fileReadBufio("./test.txt")
	//fileReadIoUtil("./test.txt")
	//fileWrite("./test.txt")
	//fileWriteBufio("./test.txt")
	fileWriteIoUtil("./test.txt")
	copyFile("./test.txt","./test2.txt")
}
