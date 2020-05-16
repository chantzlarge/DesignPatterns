package main

// Client ...
type Client struct {
	builder  *ConcreteBuilder
	director *Director
}

// NewClient ...
func NewClient() *Client {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	return &Client{
		builder:  builder,
		director: director,
	}
}
