package main

import (
	"fmt"
)

func longestPalindrome2(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	// 所有单字符都是回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 处理长度大于 1 的子串
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i == 1 { // 两个字符相等，如 "aa"
					dp[i][j] = true
				} else { // 检查内部子串是否回文
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 更新最长回文子串的起点和长度
			if dp[i][j] && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}

	return s[start : start+maxLen]
}

func main() {
	fmt.Println(longestPalindrome2("babad")) // 输出: "bab" 或 "aba"
	fmt.Println(longestPalindrome2("cbbd"))  // 输出: "bb"
}
