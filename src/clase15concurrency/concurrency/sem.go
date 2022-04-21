package main

import (
	"fmt"
	"sync"
	"time"
)

//Representation of channels
// c := make(chan int, 2) Limit the quantity of channels
// c := [][]
//Full channel
// c := [goRoutine][goRoutine]

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//Limit the quantity of go routines executing
		c <- 1
		wg.Add(1)
		//Send iterator, reference to wg and the channel
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Println("Id %d finished\n", i)
	//Free the space
	<-c
}
