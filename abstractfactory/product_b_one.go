package main

import "log"

// ProductBOne ...
type ProductBOne struct {
}

// Interact ...
func (b *ProductBOne) Interact(a AbstractProductA) {
	log.Printf("%T interacts with %T", b, a)
}
