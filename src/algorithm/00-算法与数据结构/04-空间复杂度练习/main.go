package main

import "fmt"

func max(arr []int) int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		if temp < arr[i] {
			temp = arr[i]
		}
	}
	return temp
}

func sum(arr []int) int {
	temp := 0
	for i := 0; i < len(arr); i++ {
		temp += arr[i]
	}
	return temp
}

func factorial(n int) int {
	temp := 1
	for i := 1; i <= n; i++ {
		temp *= i
	}
	return temp
}

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	hashMap := map[int]int{}
	var f func(x int) int
	f = func(x int) int {
		if hashMap[x] > 0 {
			return hashMap[x]
		}
		if x == 0 || x == 1 {
			return 1
		}else {
			result := f(x-1)+f(x-2)
			hashMap[x] = result
			return result
		}
	}
	return f(n)
}

func main() {
	//fmt.Println(max([]int{1, 3, 5, 7, 9, 10, 22}))
	//fmt.Println(sum([]int{1, 3, 5, 7, 9, 10, 22}))
	//fmt.Println(factorial(3))
	fmt.Println(Fibonacci(4))
	fmt.Println(Fibonacci2(4))
}
