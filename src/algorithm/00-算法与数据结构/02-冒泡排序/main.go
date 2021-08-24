package main

import "fmt"

// 1. 比较相邻元素，将大的放后面
// 2. 每一轮完毕，最大的数就在最后面
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

func main() {
	// 冒泡排序 时间复杂度O(n^2)
	arr := []int{3, 1, 2, 4, 5, 7, 8, 9, 22}
	fmt.Println(bubbleSort(arr))
}
