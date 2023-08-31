package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex }

func main() {
	var numPhilosophers int
	fmt.Print("Enter number of philosophers: ")
	fmt.Scan(&numPhilosophers)

	var wg sync.WaitGroup
	chopsticks := make([]Chopstick, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			left := id
			right := (id + 1) % numPhilosophers

			chopsticks[left].Lock()
			chopsticks[right].Lock()

			fmt.Printf("Philosopher %d is eating\n", id)

			chopsticks[right].Unlock()
			chopsticks[left].Unlock()

			fmt.Printf("Philosopher %d is thinking\n", id)
		}(i)
	}

	wg.Wait()
}
