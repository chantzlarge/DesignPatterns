package main

import "log"

func main() {
	creator := NewCreator()
	concreteCreator := NewConcreteCreator()

	creator.CreateProduct("")
	concreteCreator.CreateProduct("")

	product := creator.GetProduct()
	concreteProduct := concreteCreator.GetProduct()

	log.Printf("%T", product)
	log.Printf("%T", concreteProduct)
}
