package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, checkpoint chan bool, resume chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(time.Duration(id) * time.Second)

	checkpoint <- true 

	<-resume 

	fmt.Printf("Worker %d: Resumed\n", id)
}

func main() {
	numWorkers := 3
	checkpoint := make(chan bool)
	resume := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
	}


	for i := 0; i < numWorkers; i++ {
		<-checkpoint
	}

	fmt.Println("All workers have reached the checkpoint")

	
	for i := 0; i < numWorkers; i++ {
		resume <- true
	}

	wg.Wait()
	fmt.Println("All workers have resumed")
}
