package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex // 1. Use a Mutex to protect shared state
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()         // 2. Lock before modifying
	defer c.mux.Unlock() // 3. Ensure Unlock is called
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	fmt.Println("--- Mutex ---")

	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	// Spawn 1000 goroutines to increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("somekey")
		}()
	}

	wg.Wait()
	fmt.Println("Final result (should be 1000):", c.Value("somekey"))
}
