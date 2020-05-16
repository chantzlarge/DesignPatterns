package main

// ConcretePrototypeTwo implements an operation for cloning itself.
type ConcretePrototypeTwo struct{}

// NewConcretePrototypeTwo ...
func NewConcretePrototypeTwo() *ConcretePrototypeTwo {
	return &ConcretePrototypeTwo{}
}

// Clone clones itself.
func (prototype *ConcretePrototypeTwo) Clone() Prototype {
	return prototype
}
