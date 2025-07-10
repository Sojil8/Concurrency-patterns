package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i <= 6; i++ {
		out <- i
	}
	defer close(out)
}

func worker(id int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		fmt.Printf("worker %d received %d \n", id, val)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	go producer(jobs)

	wg.Wait()
	fmt.Println("All work done")

}
