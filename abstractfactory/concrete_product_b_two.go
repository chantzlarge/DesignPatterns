package main

import "log"

// ConcreteProductBTwo ...
type ConcreteProductBTwo struct {
}

// Interact ...
func (b *ConcreteProductBTwo) Interact(a AbstractProductA) {
	log.Printf("%T interacts with %T", b, a)
}
