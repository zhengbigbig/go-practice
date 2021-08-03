package calc

// 包中的标识符(变量名、函数名、结构体、接口等)如果首字母小写，表示包私有：只能在当前这个包使用
// 首字母大写的标识符可以被外部包调用
func add(x, y int) int {
	return x + y
}

func Add(x, y int) int {
	return x + y
}
