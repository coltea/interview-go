package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- true
	}()
	select {
	case <-ch:
		fmt.Println("func done")
	case <-time.After(2 * time.Second):
		fmt.Println("time out")
	}

}
