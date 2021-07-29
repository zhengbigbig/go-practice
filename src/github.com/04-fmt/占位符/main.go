package main

import "fmt"

func fmtDemo() {
	fmt.Printf("%v\n", 100)   // 打印默认类型的变量值，100
	fmt.Printf("%v\n", false) // 打印默认类型的变量值，false
	b := "你好\""
	c := struct{ name string }{"你好\""}
	fmt.Printf("%v\n", c)  // 按照该变量的默认类型打印变量值，{你好"}
	fmt.Printf("%+v\n", c) // 类似 %v，但输出结构体时会添加字段名，{name:你好"}
	fmt.Printf("%#v\n", c) // 值的 Go 语法表示，struct { name string }{name:"你好\""}
	fmt.Printf("%T\n", b)  // 打印变量类型，struct { name string }
	fmt.Printf("100%%\n")  // 打印百分号，100%

}

func intDemo() {
	d := 100
	fmt.Printf("%d\n", d)  // 打印整形的十进制 decimal 显示， 100
	fmt.Printf("%+d\n", d) // 打印整形的十进制 decimal 显示， +100
	fmt.Printf("%b\n", d)  // 打印整形的二进制 Binary 的显示， 1100100
	fmt.Printf("%c\n", d)  // 打印整形对应的 unicode 码值显示， d
	fmt.Printf("%o\n", d)  // 打印整形的八进制 octal 的显示， 144
	fmt.Printf("%#o\n", d) // 打印整形的八进制带 0 的显示， 0144
	fmt.Printf("%x\n", d)  // 打印整形的十六进制 小写 a-f 的显示， 64
	fmt.Printf("%X\n", d)  // 打印整形的十六进制 大写 A-F 的显示， 64
	fmt.Printf("%#X\n", d) // 打印整形的十六进制带 0x 的显示， 0X64
	fmt.Printf("%U\n", d)  // 打印整形的 Unicode 格式显示， U+0064
	fmt.Printf("%#U\n", d) // 打印整形的 Unicode 格式显示，带字符， U+0064 'd'
	fmt.Printf("%q\n", d)  // 打印整形的对应的单引号括起来的 go 语法字符字面值 ， 'd'

	// 指定长度的 int
	fmt.Printf("|%d|\n", d)   // 打印整形，十进制显示， |100|
	fmt.Printf("|%8d|\n", d)  // 打印整形，长度为 8 ，右对齐，左边留空（不足的用空格补齐）， |     100|
	fmt.Printf("|%08d|\n", d) // 打印整形，长度为 8 ，右对齐，不足的使用 0 补齐）， |00000100|
	fmt.Printf("|%-8d|\n", d) // 打印整形，长度为 8 ，左对齐，右边留空（不足的用空格补齐）， |100     |
}

func floatDemo() {
	f := 3.14159265487
	fmt.Printf("%b\n", f) // 打印浮点数，二进制指数的科学计数法，无小数部分， 7074237754911210p-51
	fmt.Printf("%f\n", f) // 打印浮点数，有小数部分，无指数部分，保留 6 位小数？， 3.141593
	fmt.Printf("%F\n", f) // 打印浮点数，有小数部分，无指数部分，保留 6 位小数？， 3.141593
	fmt.Printf("%e\n", f) // 打印浮点数，科学计数法显示， 3.141593e+00
	fmt.Printf("%E\n", f) // 打印浮点数，科学计数法显示， 3.141593e+00
	fmt.Printf("%g\n", f) // 打印浮点数，根据实际情况采用 %e 或 %f 格式， 3.14159265487
	fmt.Printf("%G\n", f) // 打印浮点数，根据实际情况采用 %e 或 %f 格式， 3.14159265487
}

func boolDemo() {
	fmt.Printf("%t\n", true)
}

func strDemo() {
	s := "hello世界"
	fmt.Printf("%s\n", s) // 直接输出 字符串 或者 字节 []byte
	fmt.Printf("%q\n", s) // 打印该值对应的双引号括起来的 go 语法字符串字面值
	fmt.Printf("%x\n", s) // 每个字节用两字符十六进制数表示（使用a-f）
	fmt.Printf("%X\n", s) // 每个字节用两字符十六进制数表示（使用A-F）
}

func pointerDemo() {
	p := 10
	fmt.Printf("%p\n", &p)  // 表示为十六进制，加上 0x, 0xc0000ac058
	fmt.Printf("%#p\n", &p) // 表示为十六进制，不含 0x, c0000ac058
}

func widthDemo() {
	n := 12.34
	fmt.Printf("|%f|\n", n)    // 未指定宽度，使用默认宽度，不足使用 0 补齐, |12.340000|
	fmt.Printf("|%9f|\n", n)   // 宽度 9，默认精度，不足使用 0 补齐, |12.340000|
	fmt.Printf("|%.2f|\n", n)  // 默认宽度，精度 2 两位小数, |12.34|
	fmt.Printf("|%9.2f|\n", n) // 宽度 9，精度 2，不足使用 0 补齐, |    12.34|
	fmt.Printf("|%9.f|\n", n)  // 宽度 9，精度 0 不足使用 空格 补齐, |       12|
	fmt.Printf("|%09.f|\n", n) // 宽度 9，精度 0 不足使用 0 补齐, |       12|
}

func otherDemo() {
	o := "hello世界"
	fmt.Printf("|%s|\n", o)       // 打印字符串， |hello世界|
	fmt.Printf("|%10s|\n", o)     // 打印字符串，宽度 10，默认右对齐，空格补齐， |   hello世界|
	fmt.Printf("|%-10s|\n", o)    // 打印字符串，宽度 10，负号 改为左对齐，空格补齐， |hello世界   |
	fmt.Printf("|%2s|\n", o)      // 打印字符串，宽度 2 少于字符串长度， |hello世界|
	fmt.Printf("|%-2s|\n", o)     // 打印字符串，宽度 2 少于字符串长度， |hello世界|
	fmt.Printf("|%10.10s|\n", o)  // 打印字符串，宽度 10 精度 10 ，与默认相反，右对齐，空格补齐， |   hello世界|
	fmt.Printf("|%-10.10s|\n", o) // 打印字符串，宽度 10 精度 10 ，与默认相反，左对齐，空格补齐， |hello世界   |
	fmt.Printf("|%2.2s|\n", o)    // 打印字符串，宽度 2 精度 2 少于字符串长度，保留 2 个字符，多的截掉， |he|
	fmt.Printf("|%-2.2s|\n", o)   // 打印字符串，宽度 2 精度 2 少于字符串长度，保留 2 个字符，多的截掉， |he|
	fmt.Printf("|%010s|\n", o)    // 打印字符串，宽度 10 ，精度默认，不足使用 0 补齐， |000hello世界|
}

func main() {
	fmtDemo()
	intDemo()
	floatDemo()
	boolDemo()
	strDemo()
	pointerDemo()
	widthDemo()
	otherDemo()
}