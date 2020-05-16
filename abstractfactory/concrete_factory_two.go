package main

// ConcreteFactoryTwo ...
type ConcreteFactoryTwo struct{}

// CreateProductA ...
func (factory *ConcreteFactoryTwo) CreateProductA() AbstractProductA {
	return &ProductATwo{}
}

// CreateProductB ...
func (factory *ConcreteFactoryTwo) CreateProductB() AbstractProductB {
	return &ProductBTwo{}
}
