package main

// AbstractFactory declares an interface for operations that create abstract
// product objects.
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}
