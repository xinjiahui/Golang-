package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notifyCh = make(chan struct{}, 5)

// goroutine 池

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
		notifyCh <- struct{}{}
	}

}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go func() {
		for j := 1; j <= 50; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	//开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for i := 0; i < 50; i++ {
			<-notifyCh
		}
		close(results)
	}()

	for x := range results {
		fmt.Println(x)
	}

}
