package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal1(root *TreeNode) (ret []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

// 迭代
func preorderTraversal2(root *TreeNode) (ret []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			ret = append(ret, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

// morris
func preorderTraversal3(root *TreeNode) (ret []int) {
	// 1. 有无左树，无则右移
	// 2. 有左树，找到左树最右叶子节点mostRight
	// 2.1 最右mostRight指向nil mostRight.Right = cur,cur=cur.Left
	// 2.2 最右mostRight指向cur mostRight.Right = nil,cur=cur.Right
	cur := root
	var mostRight *TreeNode = nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				ret = append(ret, cur.Val)
				mostRight.Right = cur
				cur = cur.Left
				continue
			}
			mostRight.Right = nil
		}else {
			ret = append(ret, cur.Val)
		}
		cur = cur.Right
	}
	return
}

func main() {
	node1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}
	//fmt.Println(preorderTraversal1(node1))
	fmt.Println(preorderTraversal2(node1))
	//fmt.Println(preorderTraversal3(node1))
}
