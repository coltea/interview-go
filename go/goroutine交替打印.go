package main

import (
	"fmt"
	"sync"
)

func main() {
	dogChan, catChan, fishChan := make(chan bool), make(chan bool), make(chan bool)
	var wg sync.WaitGroup
	printAni := func(ch1 chan bool, ch2 chan bool, n string, isEnd bool) {
		defer wg.Add(-1)
		i := 0
		for {
			select {
			case <- ch1:
				if i < 10 {
					fmt.Println(n, n)
					i += 1
					ch2 <- true
				} else if ! isEnd {
					ch2 <- true
					return
				} else {
					return
				}
			}
		}
	}
	wg.Add(3)
	go printAni(dogChan, catChan, "dog", false)
	go printAni(catChan, fishChan, "cat", false)
	go printAni(fishChan, dogChan, "fish", true)
	dogChan <- true
	wg.Wait()

}
