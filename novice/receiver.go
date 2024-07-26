package main

import (
	"fmt"
)

type Runner struct {
	name string
}

func (r Runner) changeName() {
	r.name = "changeName"
}

func (r Runner) Print() {
	fmt.Println(r.name)
}

func (r *Runner) changeName2() {
	r.name = "changeName2"
}

func main() {
	a := Runner{"name"}
	a.Print()
	a.changeName()
	a.Print()
	a.changeName2()
	a.Print()
}
