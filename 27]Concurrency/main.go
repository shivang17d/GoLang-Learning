package main

import (
	"fmt"
	"sync"
	"time"
)

// --- CONCEPT 1: GOROUTINES (Concurrency) ---
// Concurrency is the composition of independently executing processes.
// In Go, goroutines are extremely lightweight threads managed by the Go runtime.

func task(id int, name string) {
	fmt.Printf("Task %d (%s) started\n", id, name)
	// Simulate some work
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Task %d (%s) finished\n", id, name)
}

// --- CONCEPT 2: CHANNELS (Communication) ---
// "Do not communicate by sharing memory; instead, share memory by communicating."
// Channels allow goroutines to pass values to each other safely.

func producer(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"Hello", "from", "the", "producer"}
	for _, m := range messages {
		fmt.Printf("Producer: sending '%s'\n", m)
		ch <- m // Blocking send
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // Closing the channel signals consumers that no more data is coming
}

func consumer(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Reading from a channel until it is closed
	for msg := range ch {
		fmt.Printf("Consumer: received '%s'\n", msg)
	}
}

// --- CONCEPT 3: MUTEX (Parallelism Safety / Shared State) ---
// When multiple goroutines access shared data simultaneously (Parallelism),
// we use a Mutex (Mutual Exclusion) to prevent data races.

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()         // Lock so only one goroutine at a time can access the map
	defer c.mux.Unlock() // Ensure it unlocks even if a panic occurs
	c.v[key]++
}

func (c *SafeCounter) Value(key string) (int, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.v[key]
	return val, ok
}

func main() {
	fmt.Println("=== CONCURRENCY VS PARALLELISM IN GO ===")

	// 1. Basic Goroutines
	fmt.Println("\n--- Section 1: Basic Goroutines ---")
	go task(1, "background") // Runs task in a separate goroutine
	task(2, "foreground")    // Runs task in the main goroutine
	// Note: If we don't wait, Section 1 might disappear before background finishes!

	// 2. WaitGroups (The proper way to wait for goroutines)
	fmt.Println("\n--- Section 2: WaitGroups ---")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Track that 1 new goroutine is starting
		go func(id int) {
			defer wg.Done() // Signal that this goroutine is finished
			task(id, "worker")
		}(i)
	}
	wg.Wait() // Block main until all 3 workers call wg.Done()

	// 3. Channels (Communication)
	fmt.Println("\n--- Section 3: Channels ---")
	msgChan := make(chan string) // Unbuffered channel
	var chanWg sync.WaitGroup
	chanWg.Add(2)

	go producer(msgChan, &chanWg)
	go consumer(msgChan, &chanWg)

	chanWg.Wait()

	// 4. Mutex (Safe Shared State)
	fmt.Println("\n--- Section 4: Mutex & Shared State ---")
	counter := SafeCounter{v: make(map[string]int)}
	var mutexWg sync.WaitGroup

	// Start 1000 goroutines to increment the same counter
	for i := 0; i < 1000; i++ {
		mutexWg.Add(1)
		go func() {
			defer mutexWg.Done()
			counter.Inc("clicks")
		}()
	}
	mutexWg.Wait()
	val, _ := counter.Value("clicks")
	fmt.Printf("Final counter value: %d (Expected: 1000)\n", val)

	fmt.Println("\n=== ALL CONCURRENCY EXAMPLES FINISHED ===")
}
