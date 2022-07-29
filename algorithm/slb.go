package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Client struct {
	Name string
}

func (m *Client) Do() {
	fmt.Println("do", m.Name)
}

type LoadBalance struct {
	Client []*Client
	size   int
}

func NewLoadBalance(size int) *LoadBalance {
	LoadBalance := &LoadBalance{
		Client: make([]*Client, size),
		size:   size,
	}
	return LoadBalance
}

func (m *LoadBalance) getClient() *Client {
	rand.Seed(time.Now().Unix())
	rand.Intn(100)

}

func main() {
	lb := NewLoadBalance(6)
	lb.getClient()
}
