package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 50}
	nWorkers := 3
	//Create the channels
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	//Create the workers
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	//Iterate all the values from tasks and send through the jobs
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

//Will tell us which worker is busy
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		//Calc the serie for that number
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d", id, job, fib)
		// Send the results to the last channel
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
