package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([]int, len(s))

	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1
	out := string([]rune(s)[0])

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i-1; j++ {
			// fmt.Println(s[j:i])
			if s[j:i] == reverse(s[j:i]) {
				dp[i-1] = len(s[j:i])
				if len(s[j:i]) > maxLen {
					maxLen = len(s[j:i])
					out = s[j:i]
				}
				break
			}
		}
		// fmt.Println(dp)
	}

	return out
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(longestPalindrome("babad")) // 输出: "bab" 或 "aba"
	fmt.Println(longestPalindrome("cbbd"))  // 输出: "bb"
}
