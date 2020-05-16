package main

// ConcreteBuilder ...
type ConcreteBuilder struct {
	product *Product
}

// NewConcreteBuilder ...
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{}
}

// BuildPart ...
func (builder *ConcreteBuilder) BuildPart() {
	builder.product = NewProduct()
}

// GetResult ...
func (builder *ConcreteBuilder) GetResult() *Product {
	return builder.product
}
