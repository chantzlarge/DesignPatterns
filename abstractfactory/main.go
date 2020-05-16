package main

func main() {
	concreteFactoryOne := &ConcreteFactoryOne{}
	concreteFactoryTwo := &ConcreteFactoryTwo{}

	clientOne := New(concreteFactoryOne)
	clientTwo := New(concreteFactoryTwo)

	clientOne.Run()
	clientTwo.Run()
}
