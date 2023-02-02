package main

import (
	`fmt`
)

func main() {
	// var a map[string]int
	// a := make(map[string]int)
	// a["123"] = 1

	x := "text"
	xBytes := []byte(x)
	xBytes[0] = '1' // 注意此时的 T 是 rune 类型
	fmt.Println(len(xBytes))
	x = string(xBytes)
	fmt.Println(x) // Text

}
