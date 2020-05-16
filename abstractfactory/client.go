package main

// Client uses only interfaces declared by AbstractFactory and AbstractProduct
// classes.
type Client struct {
	ProductA AbstractProductA
	ProductB AbstractProductB
}

// New ...
func New(factory AbstractFactory) *Client {
	return &Client{
		ProductA: factory.CreateProductA(),
		ProductB: factory.CreateProductB(),
	}
}

// Run ...
func (client *Client) Run() {
	client.ProductB.Interact(client.ProductA)
}
