package main

import (
	"fmt"
	"sync"
)

func main() {

	ch1, ch2, ch3 := make(chan int, 1), make(chan int, 1), make(chan int, 1)

	wg := sync.WaitGroup{}
	wg.Add(3)
	fn := func(rev chan int, sed chan int) {
		defer wg.Done()
		for {
			select {
			case n := <-rev:
				fmt.Println(n)
				sed <- n + 1
				if n+3 > 10 {
					return
				}
			}
		}
	}
	go fn(ch1, ch2)
	go fn(ch2, ch3)
	go fn(ch3, ch1)
	ch1 <- 1
	wg.Wait()

}
