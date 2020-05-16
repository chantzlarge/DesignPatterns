package main

// AbstractProductB declares an interface for the B type of product object.
type AbstractProductB interface {
	Interact(a AbstractProductA)
}
