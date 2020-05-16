package main

// ConcreteFactoryOne ...
type ConcreteFactoryOne struct{}

// CreateProductA ...
func (factory *ConcreteFactoryOne) CreateProductA() AbstractProductA {
	return &ConcreteProductAOne{}
}

// CreateProductB ...
func (factory *ConcreteFactoryOne) CreateProductB() AbstractProductB {
	return &ConcreteProductBOne{}
}
