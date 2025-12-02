package main

type intPerson interface {
	getName() string
}

type Person struct {
	name string
}

type Account struct {
}

func (p *Person) getName() string {
	return p.name
}

func showAllCar(x interface{}) {

}

func showVar(x intPerson) {
	x.getName()
}

func main() {
	m
}
