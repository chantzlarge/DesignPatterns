package main

// Client ...
type Client struct {
	abstractProductA AbstractProductA
	abstractProductB AbstractProductB
}

// New ...
func New(factory AbstractFactory) *Client {
	return &Client{
		abstractProductA: factory.CreateProductA(),
		abstractProductB: factory.CreateProductB(),
	}
}

func (client *Client) Run() {
	client.abstractProductB.Interact(client.abstractProductA)
}
