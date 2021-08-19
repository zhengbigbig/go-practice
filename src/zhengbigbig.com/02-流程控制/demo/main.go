package main

import (
	"fmt"
)

// 一般的跳出两层循环,通过自定义标签跳出循环
func gotoDemo01() {
	var breakFlag bool // 默认的初始化值 false
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				// 设置退出标签
				breakFlag = true // 修改 breakFlag 的值为 true
				break            // 跳出 for 循环
				// continue // 继续下一次循环
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层 for 循环判断
		if breakFlag { // 判断 breakFlag 的值为 true ，执行后续语句
			fmt.Println("这是外层for循环")
			break
		}
	}
	fmt.Println("跳出两层循环-break通用语法")
}

// 使用 goto 简化跳出两层 for 循环的代码
// goto 语句通过标签进行代码间的无条件跳转
// goto 语句可以在快速跳出循环、避免重复退出上有一定的帮助
// Go 语言中使用 goto 语句能简化一些代码的实现过程,例如双层嵌套的 for 循环退出
func gotoDemo02() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag // GO 语言标签的引用
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		fmt.Println("这是外层for循环")
	}
	return
	// 标签
breakTag: // GO 语言标签的定义
	fmt.Println("跳出两层循环-goto标签语法")
}

// GO 语言中使用 break 跳出循环
// GO 语言中 break 语句可以结束 for、switch 和 select 的代码块
// GO 语言中 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块
// break 标签要求必须定义在对应的 for、switch 和 select 的代码块之上
func breakDemo() {
BREAKDEMO1: // GO语言：定义一个 break 标签，用的不多，执行到此会把该标签后的 for循环 跳出，相当于跳出两层 for 循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				break BREAKDEMO1 // GO 语言标签的引用
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		fmt.Println("这是外层for循环")
	}
	fmt.Println("跳出两层循环-break标签语法")
}

// GO 语言中使用 continue 实现继续下次循环
// continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用
// 在 continue 语句后添加标签时，表示开始标签对应的循环
func continueDemo() {
forloop1: /// GO语言：定义一个 continue 标签，用的不多，执行到此会跳过本次循环，并转到标签位置，继续执行后续 for 循环
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1 // GO 语言标签的引用
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		fmt.Println("这是外层for循环")
	}
	fmt.Println("继续循环-continue标签语法")
}

func demo() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	//gotoDemo01()
	//gotoDemo02()
	// breakDemo()
	// continueDemo()
	demo()
}
