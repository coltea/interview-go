package main

import (
	"fmt"
)

// // 测试用例
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// 示例 2:
//
// 输入: numRows = 1
// 输出: [[1]]
func main() {
	testCases := []int{1, 2, 5, 6} // 测试 numRows = 1, 2, 5, 6
	for _, numRows := range testCases {
		fmt.Printf("杨辉三角（%d 行）:\n", numRows)
		result := generate(numRows)
		for _, row := range result {
			fmt.Println(row)
		}
		fmt.Println()
	}
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	if numRows == 1 {
		return res
	}
	res[1] = []int{1, 1}
	if numRows == 2 {
		return res
	}
	for i := 2; i <= numRows-1; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j <= i-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
