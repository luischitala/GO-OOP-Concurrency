package main

import "fmt"

//Declare de object
type Employee struct {
	id   int
	name string
}

//Create a new instance
func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "Luis Ernesto Román Chitala"
	fmt.Printf("%v", e)

}
