package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 1. Leverage all available CPU cores for true parallelism
	numCores := runtime.NumCPU()
	fmt.Printf("Number of CPU cores: %d\n", numCores)
	runtime.GOMAXPROCS(numCores)

	var wg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5, 6, 7, 8} // Imagine millions of items

	// 2. Split work into parallel chunks
	chunkSize := len(data) / numCores

	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go func(chunkID int) {
			defer wg.Done()
			// This block runs at the exact same time as others on different cores
			start := chunkID * chunkSize
			fmt.Printf("Core %d processing chunk starting at index %d\n", chunkID, start)
		}(i)
	}

	wg.Wait()
	fmt.Println("All parallel tasks complete.")
}
