package main

import "fmt"

type LetterFreq map[rune]int

func CountLetter(strs []string, concurrency uint)  LetterFreq {
	letterFreq := make(LetterFreq)
	for _, str := range strs {
		for _, i := range str {
			letterFreq[i] ++
		}
	}
	return letterFreq
}

func main() {
	input := []string{"abc", "ac", "ab"}
	fmt.Println(CountLetter(input))
	//fmt.Println(CountLetter(inp ut)["a"])
	a := make(chan bool,1)
}
