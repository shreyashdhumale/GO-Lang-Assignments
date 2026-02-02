package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}
func main() {
	const (
		numWorkers = 3
		numJobs    = 10
	)
	jobs := make(chan int)
	var wg sync.WaitGroup

	// start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// close channel after sending all jobs
	close(jobs)

	// wait for all workers to finish
	wg.Wait()

	fmt.Println("All jobs processed.")
}
