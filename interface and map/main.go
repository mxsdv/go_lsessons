package main

import "fmt"

type intPerson interface {
	getName()
}

type Person struct {
	name string
}

func main() {
	var x intPerson
	x = 100
	fmt.Println(x)
}
