package main

import "log"

func main() {
	// the client creates the Director object and configures it with the
	// desired Builder object.
	client := NewClient()

	// Director notifies the builder whenever a part of the product should be
	// built.
	client.director.Construct()

	// The client retrieves the product from the builder.
	product := client.builder.GetResult()

	log.Printf("%T", product)
}
