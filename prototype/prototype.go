package main

// Prototype declares an interface for cloning itself.
type Prototype interface {
	Clone() Prototype
}
