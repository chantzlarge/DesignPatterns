package main

// ConcreteCreator ...
type ConcreteCreator struct {
	Creator
}

// NewConcreteCreator ...
func NewConcreteCreator() *ConcreteCreator {
	return &ConcreteCreator{}
}

// CreateProduct ...
func (creator *ConcreteCreator) CreateProduct(id string) {
	switch id {
	default:
		creator.product = NewConcreteProduct()
	}
}
