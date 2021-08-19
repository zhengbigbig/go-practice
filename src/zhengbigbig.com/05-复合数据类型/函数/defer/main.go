package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层分为两步执行
// 1. 返回值复制
// defer
// 2. 真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 修改的是x，不是返回值 x = 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值 = x = 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // = y = x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return 5 // 返回值 =x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

// 传一个x的指针到匿名函数中
func f6() (x int) {
	defer func(x *int) {
		*x ++
	}(&x)
	return 5
}

func f7(index string,a,b int) int  {
	ret := a+b
	fmt.Println(index,a,b,ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer f7("1",a,f7("10",a,b))
	a = 0
	defer f7("2",a,f7("20",a,b))
	b = 1
	// 1. a:= 1 b:= 2
	// 2. f7("10",a,b) "10" 1 2 3
	// 3. defer f7("1",1,3)
	// 4. a = 0
	// 5. f7("20",0,2) "20" 0 2 2
	// 6. defer f7("2",0,2)
	// 7. b = 1
	// 8. f7("2",0,2) "2" 0 2 2
	// 9. f7("1",1,3) "1" 1 3 4
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
	fmt.Println(f6()) // 6 形参传递进去都是副本，若想修改同一个值，需要传递内存地址
}
