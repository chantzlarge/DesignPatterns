package main

// ConcreteFactoryOne ...
type ConcreteFactoryOne struct{}

// CreateProductA ...
func (factory *ConcreteFactoryOne) CreateProductA() AbstractProductA {
	return &ProductAOne{}
}

// CreateProductB ...
func (factory *ConcreteFactoryOne) CreateProductB() AbstractProductB {
	return &ProductBOne{}
}
