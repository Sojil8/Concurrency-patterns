package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func (t Task) Execute(workerId int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d completed task %d\n", workerId, t.ID)
}

// Worker processes jobs from the job channel
func Worker(id int, job <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for k := range job {
		k.Execute(id)
	}
}

func main() {
	fmt.Println("Started the worker pool pattern")
	totalTasks := 10       //number of task
	numWorkers := 3        //number of worker
	var wg sync.WaitGroup  // WaitGroup to synchronize the completion of all workers
	job := make(chan Task) // Job channel to send tasks to workers

	// Launch workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, job, &wg)
	}

	// Send tasks into the job channel
	for j := 1; j <= totalTasks; j++ {
		job <- Task{ID: j}
	}

	close(job) // Close the job channel after all jobs are sent

	wg.Wait() // Wait for all workers to finish
	fmt.Println("Finish the whole work and under the workers")
}
