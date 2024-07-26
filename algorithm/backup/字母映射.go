package main

import (
	"fmt"
	"strings"
)

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}


func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digestMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v",},
		"9": {"w", "x", "y", "z"},
	}
	tmp := make([]string, 0)
	res := make([]string, 0)
	var trackBack func(index int)

	trackBack = func(index int) {
		if index == len(digits) {
			res = append(res, strings.Join(tmp, ""))
		} else {
			slt := strings.Split(digits, "")
			for _, digest := range digestMap[slt[index]] {
				tmp = append(tmp, digest)
				trackBack(index+1)
				tmp = tmp[:len(tmp)-1]
			}
		}
		return
	}

	trackBack(0)
	return res
}
