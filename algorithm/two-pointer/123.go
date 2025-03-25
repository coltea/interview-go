package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		// 跳过非字母数字字符
		for i < j && !isAlphanumeric(runes[i]) {
			i++
		}
		for i < j && !isAlphanumeric(runes[j]) {
			j--
		}

		// 比较转换为小写后的字符
		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}

		i++
		j--
	}

	return true
}
