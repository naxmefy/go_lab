package main

import "time"

type Greeter struct {
	Message Message
	Grumpy  bool
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m, Grumpy: time.Now().Unix()%2 == 0}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return "Go away!"
	}
	return g.Message
}
