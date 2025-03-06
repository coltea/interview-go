package main

import "fmt"

// 定义方向数组，表示上下左右四个方向
var directions = []struct{ x, y int }{
	{0, 1},  // 向右
	{1, 0},  // 向下
	{0, -1}, // 向左
	{-1, 0}, // 向上
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	var count int

	// 深度优先搜索 (DFS) 函数
	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 如果越界或当前格子是水，返回
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == '0' {
			return
		}
		// 将当前格子标记为已访问
		grid[x][y] = '0'
		// 遍历四个方向
		for _, dir := range directions {
			dfs(x+dir.x, y+dir.y)
		}
	}

	// 遍历网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' { // 找到一个岛屿
				count++
				dfs(i, j) // 深度优先搜索，标记整个岛屿
			}
		}
	}

	return count
}

func main() {
	// 测试用例 1
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Test Case 1 - Number of Islands:", numIslands(grid1)) // Expected Output: 1

	// 测试用例 2
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Test Case 2 - Number of Islands:", numIslands(grid2)) // Expected Output: 3

	// 测试用例 3
	grid3 := [][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
	}
	fmt.Println("Test Case 3 - Number of Islands:", numIslands(grid3)) // Expected Output: 1

	// 测试用例 4 - 只有水
	grid4 := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Println("Test Case 4 - Number of Islands:", numIslands(grid4)) // Expected Output: 1

}
