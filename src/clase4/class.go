package main

import "fmt"

//Declare de object
type Employee struct {
	id   int
	name string
}

//Receiver functions
//Setters
func (e *Employee) SetId(id int) {
	//We do not use self insted this
	e.id = id
}

func (e *Employee) SetName(name string) {
	//We do not use self insted this
	e.name = name
}

// Getters
func (e *Employee) GetId() int {
	//We do not use self insted this
	return e.id
}

func (e *Employee) GetName() string {
	//We do not use self insted this
	return e.name
}

//Create a new instance
func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "Luis Ernesto Román Chitala"
	fmt.Printf("%v", e)
	e.SetId(5)
	fmt.Printf("%v", e)
	e.SetName("El Luis")
	fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
