# Singleton

---

## Intent

Ensure a class only has one instance, and provide a global point of access to 
it.

## Structure

<img src="https://docs.google.com/drawings/d/e/2PACX-1vR4dAoljge8n4Mi9k1YUBs8dudfjzUqAR4e7vxG_-dVAm10RIedqRMieFrtfnAtfzAO0MN70t_hhlZT/pub?w=960&amp;h=720">

## Usage

Build

```
cd ./singleton
go build
./singleton
```

Example

```go
package main

import "log"

var singleton *Singleton

func init() {
	singleton = NewSingleton()
}

func main() {
	uniqueInstance := singleton.Instance()
	log.Printf("%v", uniqueInstance)

	singletonData := singleton.GetSingletonData()
	log.Printf("%v", singletonData)

	singleton.SingletonOperation()
}
```