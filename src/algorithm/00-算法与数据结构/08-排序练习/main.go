package main

import (
	"fmt"
)

// O(n^2) O(1)
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
		fmt.Printf("待排序的切片:%v", arr[i:n])
		fmt.Printf("经过 %v 次循环后的切片%v \n", i+1, arr)
	}
	return arr
}

func selectionSort2(arr []int) []int {
	n := len(arr)
	// 5 4 3 2
	for i := 0; i < n-1; i++ {
		min := i
		max := n - i - 1
		for j := i + 1; j < n-i; j++ {
			if arr[min] > arr[j] {
				min = j
			}
			if arr[max] < arr[j] {
				max = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
		arr[max], arr[n-1-i] = arr[n-1-i], arr[max]
	}
	return arr
}

// O(n^2) O(1)
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// O(n^2) O(1)
func insertSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		j := i - 1
		temp := arr[i]
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}

// O(n log n) O(1)
func shellSort(arr []int) []int {
	n := len(arr)
	h := 1
	for h <= n/3 {
		h = h*3 + 1
	}
	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < n; i++ {
			j, temp := i-gap, arr[i]
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		fmt.Printf("gap为：%v 时排序后的切片：%v\n", gap, arr)
	}
	return arr
}

// O(n log n) O(1)
func quickSort(arr []int) []int {
	quickSortPartition(arr, 0, len(arr)-1)
	return arr
}

func quickSortPartition(arr []int, left, right int) {
	if left <= right {
		return
	}
	storeIndex := partition(arr, left, right)
	quickSortPartition(arr, left, storeIndex-1)
	quickSortPartition(arr, storeIndex+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	storeIndex := left
	for i := left; i < right; i++ {
		if arr[i] <= pivot {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++
		}
	}
	arr[storeIndex], arr[right] = arr[right], arr[storeIndex]
	return storeIndex
}

// O(n log n) O(n)
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}
func merge(left, right []int) []int {
	var ret []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] > right[0] {
			ret = append(ret, right[0])
			right = right[1:]
		} else {
			ret = append(ret, left[0])
			left = left[1:]
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

// O(n+k) O(n+k)
//
func countingSort(arr []int, maxValue int) []int {
	bktLen := maxValue + 1
	bkt := make([]int, bktLen)
	n := len(arr)
	sortedIndex := 0
	for i := 0; i < n; i++ {
		bkt[arr[i]] ++
	}
	for j := 0; j < bktLen; j++ {
		for bkt[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex++
			bkt[j] --
		}
	}
	return arr
}

func radixSort(arr []int,digit int) []int {
	n := len(arr)
	l := 1
	for i := 0; i < digit; i++ {
		var bkt [10][]int
		var ret []int
		for j := 0; j < n; j++ {
			num := arr[j] / l % 10
			bkt[num] = append(bkt[num],arr[j])
		}
		for k := 0; k < 10; k++ {
			ret = append(ret,bkt[k]...)
		}
		for k := range arr {
			arr[k] = ret[k]
		}
		l *= 10
	}
	return arr
}
func main() {
	arr := []int{5, 2, 3, 1, 6, 2, 7, 9, 8}
	//fmt.Println("选择排序", selectionSort(arr))
	//fmt.Println("选择排序(优化)", selectionSort2(arr))
	//fmt.Println("冒泡排序",bubbleSort(arr))
	//fmt.Println("插入排序",insertSort(arr))
	//fmt.Println("希尔排序", shellSort(arr))
	//fmt.Println("快速排序",quickSort(arr))
	//fmt.Println("归并排序",mergeSort(arr))
	//fmt.Println("计数排序", countingSort(arr, 10))
	fmt.Println("计数排序", radixSort(arr, 1))
}
