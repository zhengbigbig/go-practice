package main

import "fmt"

// 迭代
func binarySearch(arr []int, tg int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if tg > arr[mid] {
			low = mid + 1
		} else if tg < arr[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 递归
func binarySearch2(arr []int, tg int) int {
	n := len(arr)
	start, end := 0, n-1
	return search(arr, tg, start, end)
}

func search(arr []int, tg int, start int, end int) int {
	mid := (start + end) / 2
	if end < start {
		return -1
	} else if tg > arr[mid] {
		start = mid + 1
	} else if tg < arr[mid] {
		end = mid - 1
	} else {
		return mid
	}
	return search(arr, tg, start, end)
}


func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 12, 13, 14, 52}, 5))  // 9
	fmt.Println(binarySearch2([]int{1, 2, 3, 4, 5, 6, 7, 8, 12, 13, 14, 52}, 5)) // 9
}
