# Prototype

---

## Intent

Specify the kinds of objects to create using a prototypical instance, and 
create new objects by copying this prototype.

## Structure

<img src="https://docs.google.com/drawings/d/e/2PACX-1vS9UoxTJZxoemXIECZumJDuwK9v-sdgoi-13YETQkxvYk6nurXOCkvSFKLGxWTYI7RN6RZQvGDmNNjN/pub?w=960&amp;h=720">

## Usage

Build

```
cd ./prototype
go build
./prototype
```

Example

```go
package main

import "log"

func main() {
	client := NewClient()
	prototypeOne := NewConcretePrototypeOne()
	prototypeTwo := NewConcretePrototypeTwo()

	client.Operation(prototypeOne)
	log.Printf("%T", client.p)

	client.Operation(prototypeTwo)
	log.Printf("%T", client.p)
}
```