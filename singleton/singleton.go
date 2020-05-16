package main

import "log"

// Singleton ...
type Singleton struct {
	uniqueInstance string
	singletonData  string
}

// NewSingleton ...
func NewSingleton() *Singleton {
	return &Singleton{
		uniqueInstance: "hello",
		singletonData:  "world",
	}
}

// Instance ...
func (c *Singleton) Instance() string {
	return c.uniqueInstance
}

// GetSingletonData ...
func (c *Singleton) GetSingletonData() string {
	return c.singletonData
}

// SingletonOperation ...
func (c *Singleton) SingletonOperation() {
	log.Printf("%v %v", c.uniqueInstance, c.singletonData)
}
