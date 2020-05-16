# AbstractFactory

---

## Intent

> Separate the construction of a complex object from it's representation so that the same construction process can create different representations.

## Structure

<img src="https://docs.google.com/drawings/d/e/2PACX-1vQh3dOc17AtxEJCrBD2mWvDYAVyoWww_RwmfYRXOgttxyYCA8lE2LuwQR4oAg98M7J_8LGkrOcVGeIa/pub?w=960&amp;h=720">

## Usage

```go
func main() {
	concreteFactoryOne := &ConcreteFactoryOne{}
	concreteFactoryTwo := &ConcreteFactoryTwo{}

	clientOne := New(concreteFactoryOne)
	clientTwo := New(concreteFactoryTwo)

    clientOne.Run() // output: 2020/05/16 08:10:21 *main.ConcreteProductBOne interacts with *main.ConcreteProductAOne
	clientTwo.Run() // output: 2020/05/16 08:10:21 *main.ConcreteProductBTwo interacts with *main.ConcreteProductATwo
}
```