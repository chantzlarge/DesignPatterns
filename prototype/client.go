package main

// Client ...
type Client struct {
	p Prototype
}

// NewClient ...
func NewClient() *Client {
	return &Client{}
}

// Operation ...
func (c *Client) Operation(p Prototype) {
	c.p = p.Clone()
}
