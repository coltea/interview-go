package main

import (
	"fmt"
)

func main() {
	goSelect()
}
//
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func main() {
//	c := make(chan int, 10)
//	go fibonacci(cap(c), c)
//	for i := range c {
//		fmt.Println(i)
//	}
//}

// select
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x, y = y, x+y
				print(fmt.Sprintf("x:%d y:%d  ",x,y))

		case <-quit:
				fmt.Println("quit")
				return
		}
	}
}

func goSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			print(fmt.Sprintf("  C:%d  ",<-c))
			}
		quit <- 0
	}()
	fibonacci(c, quit)
}
