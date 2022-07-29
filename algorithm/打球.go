package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main是所有Go程序的入口
func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	// 计数加2，表示要2待两个goroutine
	wg.Add(1)

	// 启动两个选手
	go player("Nick", court)
	//go player("Jack", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来
		fmt.Println(1)
		select {


		case ball := <-court:
			// 选随机数，然后用这个数来判断我们是否丢球
			n := rand.Intn(100)
			if n%5 == 0 {
				fmt.Printf("Player %s Missed\n", name)
				return
			}

			// 显示击球数，并将击球数加1
			fmt.Printf("Player %s Hit %d\n", name, ball)
			ball++

			// 将球打向对手
			court <- ball
		default:
			//ball, ok := <-court
			//if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
			//}

		}

	}
}
