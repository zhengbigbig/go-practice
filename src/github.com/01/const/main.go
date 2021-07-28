package main

//const pi = 3.1415
//const e = 2.7182
// 声明这两个常量后，在整个程序运行期间它们的值都不能再发生变化了
// 多个常量一起声明
const (
	pi = 3.1415
	e = 2.7182
)
// const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

func main() {
	const a = 1
}
