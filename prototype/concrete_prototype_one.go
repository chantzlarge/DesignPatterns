package main

// ConcretePrototypeOne implements an operation for cloning itself.
type ConcretePrototypeOne struct{}

// NewConcretePrototypeOne ...
func NewConcretePrototypeOne() *ConcretePrototypeOne {
	return &ConcretePrototypeOne{}
}

// Clone clones itself.
func (prototype *ConcretePrototypeOne) Clone() Prototype {
	return prototype
}
