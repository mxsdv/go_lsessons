package main

import "fmt"

type intPerson interface {
	getName() string
}

type Person struct {
	name string
}

func (p *Person) getName() string {
	return p.name
}

func main() {
	var x intPerson
	var y Person
	x = &y
	fmt.Println(x)
}
