package main

import "fmt"

func main() {
	// Create a channel that transfers int without buffer

	// c := make(chan int)
	// c <- 1

	// fmt.Println(<-c)

	//Channel with buffer

	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	//Exceed the capacity of buffers
	// c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//Exceed the capacity of buffers
	// fmt.Println(<-c)
}
