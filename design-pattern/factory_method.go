package main

import "fmt"

// Product
type Product interface {
	Use()
}

// Concrete Product 1
type ConcreteProductA struct{}

func (p ConcreteProductA) Use() {
	fmt.Println("Using Concrete Product A")
}

// Concrete Product 2
type ConcreteProductB struct{}

func (p ConcreteProductB) Use() {
	fmt.Println("Using Concrete Product B")
}

// Factory Interface
type Factory interface {
	CreateProduct() Product
}

// Concrete Factory 1
type ConcreteFactoryA struct{}

func (f ConcreteFactoryA) CreateProduct() Product {
	return ConcreteProductA{}
}

// Concrete Factory 2
type ConcreteFactoryB struct{}

func (f ConcreteFactoryB) CreateProduct() Product {
	return ConcreteProductB{}
}

func main() {
	factoryA := ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	productA.Use()

	factoryB := ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	productB.Use()
}
