package main

// 打开文件写内容

func main() {
	//file,err := os.OpenFile("./xxx.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	//if err != nil {
	//	fmt.Println("open file failed,err:",err)
	//	return
	//}
	//defer file.Close()
	//str := "hello"
	//file.Write([]byte(str)) // 写入字节切片数据
	//file.WriteString("hello jzz") // 直接写入字符串数据

	//file,err := os.OpenFile("./xxx.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666)
	//if err != nil {
	//	fmt.Println("open file failed,err:",err)
	//	return
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(file)
	//for i := 0; i < 10; i++ {
	//	writer.WriteString("hello\n") // 将数据先写入缓存
	//}
	//writer.Flush() // 将缓存中的内容写入文件
	//str := "hello"
	//err := ioutil.WriteFile("./xxx.txt",[]byte(str),0666)
	//if err != nil {
	//	fmt.Println("write file failed,err:",err)
	//	return
	//}

}
