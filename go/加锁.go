package main

import (
	"fmt"
	"sync"
)

func main()  {
	var lo sync.Mutex

	total := 0
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		lo.Unlock()
		fmt.Println("sum:%d i %d", sum, i)

		go func () {
			total += i
			fmt.Println("total:%d i %d", total, i)

			lo.Lock()
		}()
	}
	fmt.Println("total:%d sum %d", total, sum)
}
