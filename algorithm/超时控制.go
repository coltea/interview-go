package main

import (
	"fmt"
	"time"
)

//func main() {
//	ch := make(chan struct{}, 1)
//	go func() {
//		fmt.Println("do something...")
//		time.Sleep(4 * time.Second)
//		ch <- struct{}{}
//	}()
//
//	select {
//	case <-ch:
//		fmt.Println("done")
//	case <-time.After(3 * time.Second):
//		fmt.Println("timeout")
//	}
//}

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
