package main

import "fmt"

// chan<- to write
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	//Function to close a channel
	close(c)
}

//Function to create an input channel and an outpot channel int both
//to read <-chan

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
		//Exceed the go subroutines
		// in <- 1
	}
	//Close the output channel
	close(out)
}

//Function to print the received values from the int type channel
func Print(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)
	//Read
	go Generator(generator)
	//Write
	go Double(generator, doubles)
	Print(doubles)
}
