package main

// Director constructs objects using the Building interface.
type Director struct {
	builder Builder
}

// NewDirector constructs a new Director.
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// Construct an object using the Builder interface.
func (director *Director) Construct() {
	// Builder handles requests from the director and adds parts to the
	// product.
	director.builder.BuildPart()
}
