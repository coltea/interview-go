package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
}

func lengthOfLongestSubstring(s string) int {
	t := []rune("")
	res := 0
	for i := range s {
		// fmt.Println(string(t), " ", i)
		if !slices.Contains(t, rune(s[i])) {
			t = append(t, rune(s[i]))
		} else {
			index := slices.Index(t, rune(s[i]))
			t = append(t[index+1:], rune(s[i]))
		}

		if len(t) > res {
			res = len(t)
		}
	}
	return res
}
