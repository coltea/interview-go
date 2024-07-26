package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
}

func (*Animal) Eat(str1 string) {
	fmt.Printf("eat %s", str1)
}

func main() {
	ani := Animal{}
	f := reflect.ValueOf(&ani).MethodByName("Eat")
	f.Call([]reflect.Value{reflect.ValueOf("apple")})
}
