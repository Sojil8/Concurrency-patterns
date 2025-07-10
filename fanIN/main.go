package main

import (
	"fmt"
	"sync"
)

func producer(id int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 3; i++ {
		out <- id*10 + i
	}
}

func fanIn(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, c := range cs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go producer(1, c1, &wg)
	go producer(2, c2, &wg)

	go func() {
		wg.Wait()
		close(c1)
		close(c2)
	}()

	result := fanIn(c1, c2)

	for v := range result {
		fmt.Println("Received:", v)
	}

}
