package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
}

type Human struct {
	*Person
	Power int
}

func main() {
	human := &Human{
		Person: &Person{"Archana"},
		Power:  9000,
	}

	human.Introduce()
}
