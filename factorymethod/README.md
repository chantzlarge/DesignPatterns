# Factory Method

---

## Intent

Define an interface for creating an object, but let subclasses decide which 
class to instantiate. Factory Method lets a calss defer instantiation to subclasses.

## Structure

<img src="https://docs.google.com/drawings/d/e/2PACX-1vS9UoxTJZxoemXIECZumJDuwK9v-sdgoi-13YETQkxvYk6nurXOCkvSFKLGxWTYI7RN6RZQvGDmNNjN/pub?w=960&amp;h=720">

## Usage

Build

```
cd ./factorymethod
go build
./factorymethod
```

Example

```go
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
```