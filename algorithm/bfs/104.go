package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathLength(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)   // 左子树高度
		right := dfs(node.Right) // 右子树高度
		maxDiameter = max(maxDiameter, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil},
	}
	fmt.Println("最大通路长度:", maxPathLength(root)) // 输出: 3
}
