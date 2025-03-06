package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	end := make(chan bool)

	f := func(in chan int, out chan int, name string) {
		for {
			select {
			case num := <-in:
				fmt.Println(name, num)
				if num == 10 {
					end <- true
					return
				} else {
					num++
					out <- num
				}
			}
		}
	}

	go f(a, b, "goroutine a :")
	go f(b, c, "goroutine b :")
	go f(c, a, "goroutine c :")

	a <- 1

	select {
	case <-end:
		return
	}
}
