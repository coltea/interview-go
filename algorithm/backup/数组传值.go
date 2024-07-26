package main

import (
	"fmt"
	"runtime"
)

func main() {

	// map
	persons := make(map[string]int)
	persons["张三"] = 19
	mp := &persons
	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modify(persons)
	fmt.Println("map值被修改了，新值为:", persons)

	// struct
	fmt.Println("------struct传递------")
	// p:=Person{"张三"}
	// fmt.Printf("原始Person的内存地址是：%p\n",&p)
	// modify2(p)
	// fmt.Println(p)

	// slice传递
	fmt.Println("------slice传递------")
	ages := []int{6, 6, 6}
	fmt.Printf("原始slice的内存地址是%p\n", ages)
	modify3(ages)
	fmt.Println("ages:", ages)

	fmt.Println("number of goroutines:", runtime.NumGoroutine())
	// fmt.Println("number of goroutines:", runtime.GOMAXPROCS())
}

func modify3(ages []int) {
	fmt.Printf("函数里接收到slice的内存地址是%p\n", ages)
	ages[0] = 1
	ages = append(ages, 8)
	fmt.Println("ages append:", ages)
}

func modify(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}

//
// type Person struct {
//	Name string
// }
//
// func modify2(p Person) {
//	fmt.Printf("函数里接收到Person的内存地址是：%p\n",&p)
//	p.Name = "李四"
// }
