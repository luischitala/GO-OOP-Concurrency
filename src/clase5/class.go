package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//Formula 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	//Formula 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	//Formula 3
	e3 := new(Employee)

	fmt.Printf("%v\n", *e3)

	e3.id = 1
	e3.name = "Name"
	fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(1, "Nombre", false)
	fmt.Printf("%v\n", *e4)

}
