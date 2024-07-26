package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	a := make([]Person, 0)
	b := new(map[string]int)
	fmt.Println(a, b)
}
