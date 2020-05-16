package main

// ConcreteFactoryTwo ...
type ConcreteFactoryTwo struct{}

// CreateProductA ...
func (factory *ConcreteFactoryTwo) CreateProductA() AbstractProductA {
	return &ConcreteProductATwo{}
}

// CreateProductB ...
func (factory *ConcreteFactoryTwo) CreateProductB() AbstractProductB {
	return &ConcreteProductBTwo{}
}
