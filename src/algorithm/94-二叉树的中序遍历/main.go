package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// morris遍历，打印时机：
// 1. 前序，第二次到达不打印
// 2. 中序，有左树第一次不打印，打印有左树第二次和无左树
// 3. 后续，第二次到达左树右边界逆序打印，所有过程完成后单独打印整树右边界

// 递归
func inorderTraversal1(root *TreeNode) []int {
	var ret []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		ret = append(ret, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	return ret
}

// 前序-迭代 显示模拟栈
func inorderTraversal2(root *TreeNode) (ret []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		ret = append(ret,root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return ret
}

// 前序-morris
// 1. 设置当前指针为root
// 2. 判断current无左树，cur=cur.right
// 3. cur有左树，找到左树最右节点mostRight
// 3.1 mostRight的右指针若指向nil，mostRight.right=cur,cur=cur.left
// 3.2 mostRight的右指针若指向cur，mostRight.right=nil,cur=cur.right
func inorderTraversal3(root *TreeNode) []int {
	var ret []int
	var p, q *TreeNode = root, nil
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
			ret = append(ret, p.Val)
			q.Right = nil
		}else {
			ret = append(ret, p.Val)
		}
		p = p.Right
	}
	return ret
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
	fmt.Println(inorderTraversal2(node1))
}
