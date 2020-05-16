package main

// AbstractFactory ...
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}
