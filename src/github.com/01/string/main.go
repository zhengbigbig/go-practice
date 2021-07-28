package main

import (
	"fmt"
	"strings"
)

func main() {
	// 使用 双引号 定义 单行字符串 变量
	s1 := "hello" //定义英文字符串
	s2 := "你好"    // 可以识别非 ASCII 码，默认支持 utf-8 编码

	// 使用 反引号 定义 多行字符串 变量
	s3 := `
11111
22222   // 多行字符串里的，注释不生效
33333
`

	// 使用 单引号 定义 单个字符 变量
	s4 := 'h'

	fmt.Println("\n单行字符串01", s1)
	fmt.Println("\n单行字符串02", s2)
	fmt.Println("\n多行字符串", s3)
	fmt.Println("\n字符定义", s4)

	// 字符串操作
	// 字符串求长度
	s5 := "zhongguojueqi"
	fmt.Println("\n字符串-求长度", len(s5))

	// 字符串拼接
	s6 := "nihaoshijie"
	fmt.Println("\n字符串-拼接01", s5+s6)

	s7 := fmt.Sprintf("%s---%s", s5, s6)
	fmt.Println("\n字符串-拼接02", s7)

	// 字符串分割
	s8 := strings.Split(s6, "o")
	fmt.Println("\n字符串-分割", s8)

	// 字符串包含判断
	s9 := strings.Contains(s5, "o")
	fmt.Println("\n字符串-包含判断01", s9)
	fmt.Println("\n字符串-包含判断02", strings.Contains(s5, "o"))

	// 字符串前缀， 后缀判断
	fmt.Println("\n字符串-前缀判断", strings.HasPrefix(s5, "zhong"))
	fmt.Println("\n字符串-后缀判断", strings.HasSuffix(s5, "qi"))

	// 字符串索引查找
	fmt.Println("\n字符串-索引查找-第一个字符 o 的索引", strings.Index(s5, "o"))
	fmt.Println("\n字符串-索引查找-最后一个字符 o 的索引", strings.LastIndex(s5, "o"))

	// 字符串-join操作
	s10 := []string{"aaa", "bbb", "ccc", "ddd"}
	fmt.Println("\n字符串-join 操作", strings.Join(s10, " + "))
}