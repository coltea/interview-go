package main

import (
	"fmt"
)

// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(isSymmetricBfs(&TreeNode{
		1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}))
	fmt.Println(isSymmetricBfs(&TreeNode{
		1, nil, nil}))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func isSymmetricBfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left, node2.Right)
		queue = append(queue, node1.Right, node2.Left)
	}
	return true
}
