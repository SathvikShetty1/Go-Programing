package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var num int
	fmt.Println("Enter the number of workers: ")
	fmt.Scan(&num)
	start := make(chan struct{})
	done := make(chan struct{})

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-start
			fmt.Printf("Worker %d reached the checkpoint\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d completed its work\n", id)
		}(i)
	}

	fmt.Println("Starting workers...")
	close(start)

	wg.Wait()

	close(done)
	fmt.Println("All Workers completed their work...")
}
