package main

import (
	"fmt"
	"math"
)

func main() {
	// 默认Go语言中的小数都是float64类型
	f1 := 1.2345
	f2 := float32(1.2345)
	fmt.Printf("%f \n", f1)
	fmt.Printf("%.2f \n", f2)
	fmt.Printf("%T \n", 1.23455)
	fmt.Printf("%f \n",float32(.232))
	fmt.Printf("%f \n",math.MaxFloat32)
	fmt.Printf("%f \n",math.MaxFloat64)
	f1 = f2  // float32类型的值不能直接赋值给float64类型的变量
}
