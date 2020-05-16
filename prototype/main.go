package main

import "log"

func main() {
	client := NewClient()

	prototypeOne := NewConcretePrototypeOne()
	prototypeTwo := NewConcretePrototypeTwo()

	client.Operation(prototypeOne)
	log.Printf("%T", client.p) // 2020/05/16 13:14:42 *main.ConcretePrototypeOne

	client.Operation(prototypeTwo)
	log.Printf("%T", client.p) // 2020/05/16 13:14:42 *main.ConcretePrototypeTwo
}
