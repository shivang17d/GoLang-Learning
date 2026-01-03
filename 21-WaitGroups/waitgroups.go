package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// 2. Defer wg.Done() to ensure it's called even if the function returns early.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Simulate work
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("--- WaitGroups ---")

	// 1. Create a WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// 3. Increment the WaitGroup counter
		wg.Add(1)

		// 4. Important: Pass the address of wg (&wg) to the worker
		go worker(i, &wg)
	}

	// 5. Wait for all workers to finish
	fmt.Println("Main: waiting for workers...")
	wg.Wait()
	fmt.Println("Main: all workers finished.")
}
