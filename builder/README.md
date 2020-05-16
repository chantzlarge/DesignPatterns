# Builder

---

## Intent

Separate the construction of a complex object from its representation so that 
the same construction process can create differen representations.

## Structure

<img src="https://docs.google.com/drawings/d/e/2PACX-1vS2h138rKC55G79DAm-yWIwizsY8_d1uSMVgUnizu6PP5uZclO5ALTrxzOt3GeKtBH2MiUJnQrevZKi/pub?w=480&amp;h=360">

## Usage

## Building from CLI

```bash
go build
./builder
```

## Example

```go
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
```

