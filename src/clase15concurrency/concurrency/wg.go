package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//Create wait group

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	//Stop the program until the counter stops
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("End")

}
