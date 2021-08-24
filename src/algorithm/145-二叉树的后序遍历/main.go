package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func postorderTraversal1(root *TreeNode) (ret []int) {
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ret = append(ret, node.Val)
	}
	postorder(root)
	return
}

// 迭代
func postorderTraversal2(root *TreeNode) (ret []int) {
	var stack []*TreeNode
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			ret = append(ret,root.Val)
			pre = root
			root = nil
		}else {
			stack = append(stack,root)
			root = root.Right
		}
	}
	return
}

// morris
func postorderTraversal3(root *TreeNode) (ret []int) {
	var p,q *TreeNode = root,nil
	for p != nil {
		q = p.Left
		if q != nil {
			for q.Right != nil && q.Right != p {
				q = q.Right
			}
			if q.Right == nil {
				q.Right = p
				p = p.Left
				continue
			}
			q.Right = nil
			ret = printEnd(p.Left,ret)
		}
		p = p.Right
	}
	ret = printEnd(root,ret)
	return
}

func printEnd(root *TreeNode,ret []int) []int  {
	tail := turnOffNode(root)
	cur:= tail
	for cur != nil {
		fmt.Println(ret,cur.Val)
		ret = append(ret,cur.Val)
		cur = cur.Right
	}
	turnOffNode(tail)
	return ret
}

func turnOffNode(node *TreeNode) *TreeNode {
	var beg *TreeNode = nil
	mid,end := node,node.Right
	for {
		mid.Right = beg
		if end == nil {
			break
		}
		beg,mid,end = mid,end,end.Right
	}
	return mid
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
	fmt.Println(postorderTraversal3(node1)) // 35421
}
