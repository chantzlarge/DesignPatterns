package main

// ConcreteProduct ...
type ConcreteProduct struct {
	Product
}

// NewConcreteProduct ...
func NewConcreteProduct() *ConcreteProduct {
	return &ConcreteProduct{}
}
