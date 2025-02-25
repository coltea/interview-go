package main

import (
	"fmt"
)

func main() {
	m := map[string]interface{}{
		"a": 0,
	}
	change(m)
	fmt.Println(m)

	s := []int{0}
	changeSlice(s)
	fmt.Println(s)

	s1 := []int{0}
	changeSlice1(s1)
	fmt.Println(s1)
}

func change(m map[string]interface{}) {
	m["a"] = 1
}

func changeSlice(s []int) {
	s = append(s, 1)
}

func changeSlice1(s []int) {
	s[0] = 1
}
