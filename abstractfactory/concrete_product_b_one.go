package main

import "log"

// ConcreteProductBOne ...
type ConcreteProductBOne struct {
}

// Interact ...
func (b *ConcreteProductBOne) Interact(a AbstractProductA) {
	log.Printf("%T interacts with %T", b, a)
}
