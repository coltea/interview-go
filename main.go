package main

import (
	"fmt"

	"go.uber.org/fx"
)

type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{B: b}
}

type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{c}
}

type C struct {
}

func NewC() *C {
	return &C{}
}
func main() {
	fx.New(
		fx.Provide(NewB),
		fx.Provide(NewA),
		fx.Provide(NewC),
		fx.Invoke(PrintA),
	)
}
func PrintA(a *A) {
	fmt.Println(*a)
}
