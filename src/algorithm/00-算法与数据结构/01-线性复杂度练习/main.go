package main

import "fmt"

// 求数组、链表的最大值
func max(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}

// 寻找数组中的重复元素
func findArrRepeatElements(nums []int) []int {
	hashMap := make(map[int]int, len(nums))
	var ret []int
	for i := 0; i < len(nums); i++ {
		if hashMap[nums[i]] == 1 {
			ret = append(ret, nums[i])
			hashMap[nums[i]]++
		} else {
			hashMap[nums[i]]++
		}
	}
	return ret
}

// 判断链表是否存在环
type node struct {
	Val  int
	Next *node
}

func isLoop(head *node) bool {
	// 定义两个指针，p走一步，q走两步
	p, q := head, head
	for {
		p = p.Next
		if p == nil {
			return false
		}
		q = q.Next.Next
		if q == nil {
			return false
		}
		if p == q {
			return true
		}
	}
}

// Factorial 求阶层
func Factorial(n uint64) (ret uint64) {
	if n > 0 {
		ret = n * Factorial(n-1)
		return
	}
	return 1
}

// 翻转链表-迭代
func overTurn(head *node) (ret *node) {
	var beg *node = nil
	mid, end := head, head.Next
	for {
		mid.Next = beg
		if end == nil {
			break
		}
		// 指针整体向后移动
		beg, mid, end = mid, end, end.Next
	}
	ret = mid
	return ret
}
// 翻转链表-递归
func overTurn2(head *node) *node {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := overTurn2(head.Next)
	// 将下一个节点翻转
	head.Next.Next = head
	head.Next = nil
	return newNode
}
// 翻转链表-头插法

func main() {
	//fmt.Println(max([]int{1, 3, 5, 7, 9, 12, 23}))
	//fmt.Println(findArrRepeatElements([]int{1, 2, 3, 4, 5, 1, 1, 2, 3, 4}))
	//node1 := &node{
	//	Val: 111,
	//}
	//node2 := &node{
	//	Val:222,
	//}
	//node3 := &node{
	//	Val: 333,
	//}
	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = node1
	//fmt.Println(isLoop(node1))
	//fmt.Println(Factorial(4))
	node1 := &node{
		Val: 111,
	}
	node2 := &node{
		Val: 222,
	}
	node3 := &node{
		Val: 333,
	}
	node1.Next = node2
	node2.Next = node3
	n1 := overTurn2(node1)
	for n1 != nil {
		fmt.Println(n1.Val)
		n1 = n1.Next
	}
}
