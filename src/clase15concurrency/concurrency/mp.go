//Multiplex messages
package main

import (
	"fmt"
	"time"
)

func main() {
	//Create two channels
	c1 := make(chan int)
	c2 := make(chan int)
	//Defeer from duration
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	for i := 0; i < 2; i++ {
		//Allows to have different cases to work with different channels to know which is the channel that has been used
		select {
		// In the case that this message arrives
		case channelMsg1 := <-c1:
			//
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			//
			fmt.Println(channelMsg2)

		}

	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
