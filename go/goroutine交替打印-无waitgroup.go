package main

import (
	`fmt`
)

func main() {
	endChan := make(chan bool)
	a, b, c := make(chan bool), make(chan bool), make(chan bool)

	go chanPrint(endChan, a, b, "a")
	go chanPrint(endChan, b, c, "b")
	go chanPrint(endChan, c, a, "c")

	a <- true
	for {
		select {
		case <-endChan:
			return
		}
	}
}

func chanPrint(endChan chan bool, inChan chan bool, outChan chan bool, name string) {
	n := 0
	for {
		select {
		case <-inChan:
			fmt.Println(name, n)
			n++
			if n == 10 {
				if name != "c" {
					outChan <- true
				}
				endChan <- true
				return
			}
			outChan <- true
		}
	}
}
