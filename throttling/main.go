package main

import (
	"fmt"
	"time"
)

func ThrottledWorker(id int,in <-chan int,delay time.Duration){
	for job:=range in{
		fmt.Printf("Worker %d started job %d at %s\n", id, job, time.Now().Format("15:04:05.000"))
		time.Sleep(delay)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main(){
	jobs:=make(chan int)

	go ThrottledWorker(1,jobs,500*time.Millisecond)
	for i:=0;i<=5;i++{
		jobs<-i
	}
	close(jobs)



}

