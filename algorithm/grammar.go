package main

import "fmt"


func main() {
	//pointer()
	//structT()
	sliceLenCap()
}

// 指针
func pointer() {
	i, j := 42, 2701

	p := &i         // 指向 i

	fmt.Println(p)
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

// 结构体
type Vertex struct {
	X int
	Y int
}

func structT() {
	//v := Vertex{1, 2}
	//p := &v
	//p2 := &v
	//fmt.Println(p2)
	//
	//p.X = p.X*200
	//fmt.Println(v)
	//fmt.Println(p2)

	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p3  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)

	fmt.Println(v1, v2, v3, p3)
}


func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}


func sliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}