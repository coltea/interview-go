package main

import (
	"fmt"
)

 func printAni(ch1 chan bool, ch2 chan bool, n string, isEnd bool) {
	i := 0
	for {
		select {
		case <-ch1:
			if i < 10 {
				fmt.Println(n)
				i += 1
				ch2 <- true
			} else if ! isEnd {
				ch2 <- true
				return
			} else {
				endChan <- true
				return
			}
		}
	}
}

func main() {
	dogChan, catChan, fishChan := make(chan bool), make(chan bool), make(chan bool)
	endChan := make(chan bool)

	go printAni(dogChan, catChan, "dog", false)
	go printAni(catChan, fishChan, "cat", false)
	go printAni(fishChan, dogChan, "fish", true)
	dogChan <- true
	select {
	case <-endChan:
		return
	}

}
