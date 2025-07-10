package main

import "fmt"

func checkOdd(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			if v%2 == 0 {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for ch := range in {
			out <- ch * ch

		}
		close(out)
	}()
	return out
}

func main() {
	ch := checkOdd(7, 4, 56, 7, 4, 3)
	ch1 := sq(ch)

	for r := range ch1 {
		fmt.Println(r)
	}
}
