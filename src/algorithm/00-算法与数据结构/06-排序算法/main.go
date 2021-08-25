package main

import "fmt"

// 1. 冒泡排序
// 复杂度：平均时间复杂度O(n^2)，空间复杂度O(1) 稳定
// 核心：最大值后移
// 进行一次扫描，每次扫描再扫描对比次数进行交换
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 2.1 快速排序 （递归）
// 复杂度：平均时间复杂度O(n log n) 空间复杂度O(log n)
// 核心：每次将数据划分为两部分，合并递归，需要判断长度1退出
// 递归实现，需要额外的辅助空间
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		p := arr[len(arr)/2]
		var (
			left  []int
			right []int
		)
		for _, v := range arr {
			if v >= p {
				right = append(right, v)
			} else {
				left = append(left, v)
			}
		}
		return append(append(quickSort(left),p),quickSort(right)...)
	}
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 11, 13, 12, 22, 33, 55, 44}
	//fmt.Println(bubbleSort(arr))
	fmt.Println(quickSort(arr))
}
