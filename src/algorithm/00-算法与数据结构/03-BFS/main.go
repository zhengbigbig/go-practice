package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) {
	var lst []*TreeNode
	lst = append(lst, root)
	for len(lst) > 0 {
		cur := lst[0]
		lst = lst[1:]
		if cur.Left != nil {
			lst = append(lst, cur.Left)
		}
		if cur.Right != nil {
			lst = append(lst, cur.Right)
		}
		fmt.Println(cur.Val)
	}
}

func main() {
	node1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	bfs(node1)
}
