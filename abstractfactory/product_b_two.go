package main

import "log"

// ProductBTwo ...
type ProductBTwo struct {
}

// Interact ...
func (b *ProductBTwo) Interact(a AbstractProductA) {
	log.Printf("%T interacts with %T", b, a)
}
