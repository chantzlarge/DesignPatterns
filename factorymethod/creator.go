package main

import "log"

// Creator ...
type Creator struct {
	product Product
}

// NewCreator ...
func NewCreator() *Creator {
	return &Creator{}
}

// CreateProduct ...
func (creator *Creator) CreateProduct(id string) {
	log.Printf("not implemented")
}

// GetProduct ...
func (creator *Creator) GetProduct() Product {
	return creator.product
}
