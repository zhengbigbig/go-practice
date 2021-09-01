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

// 2. 选择排序
// 复杂度：平均时间复杂度O(n^2)，空间复杂度O(1) 不稳定
// 核心：选择出最小/大放置首/尾
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 3. 插入排序
// 复杂度：平均时间复杂度O(n^2)，空间复杂度O(1) 不稳定
// 核心：第一轮扫描，第二轮根据当前值比较之前的序列，大于当前值则向后赋值，指针往前走，直至指针走出序列
func insertionSort(arr []int) []int {
	for i := range arr {
		cur := arr[i]
		preIndex := i - 1
		for preIndex >= 0 && arr[preIndex] > cur {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = cur
	}
	return arr
}

// 4.1 快速排序 （递归）
// 复杂度：平均时间复杂度O(n log n) 空间复杂度O(n) 不稳定
// 核心：每次将数据划分为两部分，合并递归，需要判断长度1退出
// 递归实现，需要额外的辅助空间
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		p := arr[len(arr)/2]
		var left []int
		var right []int
		for _, v := range arr {
			if v > p {
				right = append(right, v)
			} else if v < p {
				left = append(left, v)
			}
		}
		return append(append(quickSort(left), p), quickSort(right)...)
	}
}

// 4.2 快速排序 （原地分割 in-place）
// 复杂度：平均时间复杂度O(n log n) 空间复杂度O(log n)
// 核心：双指针，指向待交换位置和基准值，以最右侧为基准值（改进需要将基准值移动至最后又移动回来）

func quickSort2(arr []int) []int {
	quickSortPartition(arr, 0, len(arr)-1)
	return arr
}
func quickSortPartition(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(arr, left, right)
	quickSortPartition(arr, left, pivotIndex-1)
	quickSortPartition(arr, pivotIndex+1, right)
}

// 大于基准值，指针停止，等待小于基准值，交换，指针右移
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	storeIndex := left
	for i := left; i < right; i++ {
		if arr[i] <= pivot {
			swap(arr, i, storeIndex)
			storeIndex++
		}
	}
	swap(arr, right, storeIndex)
	return storeIndex
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

// 5. 归并排序
// 复杂度：平均时间复杂度O(n log n) 空间复杂度O(n) 稳定
// 核心：分治法，合并两边有序返回
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left, right := arr[0:mid], arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var ret []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		ret = append(ret, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		ret = append(ret, right[0])
		right = right[1:]
	}
	return ret
}

// 6. 希尔排序
// 复杂度：平均时间复杂度O(n log n) 空间复杂度O(1) 不稳定
// 核心：先分组，再插入排序，重复到1
func shellSort(arr []int) []int {
	gap := len(arr) / 2
	for gap >= 1 {
		for i := gap; i < len(arr); i++ {
			cur := arr[i]
			j := i-gap
			for j>=0 && arr[j] > cur {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = cur
		}
		gap /= 2
	}
	return arr
}

// 7. 堆排序
// 复杂度：平均时间复杂度O(n log n) 空间复杂度O(1) 不稳定

// 8. 桶排序
// 复杂度：平均时间复杂度O(n+k) 空间复杂度O(n+k) 稳定

// 9. 基数排序
// 复杂度：平均时间复杂度O(n+k) 空间复杂度O(n+k) 稳定
// 核心：数组下标为值，值为个数

// 10. 计数排序
// 复杂度：平均时间复杂度O(n+k) 空间复杂度O(n+k) 稳定
func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始化为0数组
	n := len(arr)
	sortedIndex := 0
	for i := 0; i < n; i++ {
		bucket[arr[i]] ++
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex++
			bucket[j] --
		}
	}
	return arr
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 11, 13, 12, 22, 33, 55, 44}
	//fmt.Println(bubbleSort(arr))
	//fmt.Println(selectionSort(arr))
	//fmt.Println(insertionSort(arr))
	//fmt.Println(quickSort(arr))
	//fmt.Println(quickSort2(arr))
	//fmt.Println(mergeSort(arr))
	fmt.Println(shellSort(arr))
}
