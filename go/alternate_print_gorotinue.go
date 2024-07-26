package main

import (
	"fmt"
)

func main() {

	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	chClose := make(chan bool)
	fn := func(rev chan int, sed chan int) {
		for {
			select {
			case n := <-rev:
				fmt.Println(n)
				if n == 10 {
					chClose <- true
				} else {
					sed <- n + 1
				}
				if n >= 8 {
					return
				}
			}
		}
	}
	go fn(ch1, ch2)
	go fn(ch2, ch3)
	go fn(ch3, ch1)
	ch1 <- 1
	select {
	case <-chClose:
		return
	}
}
