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
