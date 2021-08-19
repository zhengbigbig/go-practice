package main

import "fmt"

// 复杂度分析
// 时间复杂度：O(N^2)，其中 N 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
// 空间复杂度：O(1)。
// [1,2,3,4]
// 思路：暴力遍历，只使用一次

func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 复杂度分析
// 时间复杂度：O(N).对每一个元素x，我们可以O(1)地寻找target - x
// 空间复杂度: O(N)，其中N是数组中的元素数量。主要为哈希表的开销
// 思路：将值存hash的key，索引存hash value，遍历看差是否在hash

func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(twoSum2(a[:], 5))
}