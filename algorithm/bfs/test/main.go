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
	fmt.Println(isSymmetricDfs(&TreeNode{
		1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}))
	fmt.Println(isSymmetricBfs(&TreeNode{
		1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}))
}

func isSymmetricDfs(node *TreeNode) bool {
	if node == nil {
		return true
	}
	return Dfs(node.Left, node.Right)
}

func Dfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if right.Val != left.Val {
		return false
	}
	return Dfs(left.Left, right.Right) && Dfs(left.Right, right.Left)
}

func isSymmetricBfs(node *TreeNode) bool {
	if node == nil {
		return true
	}
	queue := []*TreeNode{node.Left, node.Right}
	for len(queue) > 0 {
		n1, n2 := queue[0], queue[1]
		queue = queue[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}
		queue = append(queue, n1.Left, n2.Right, n1.Right, n2.Left)
	}

	return true
}
