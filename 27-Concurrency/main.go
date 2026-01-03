package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// makeBreakfast simulates a breakfast item being prepared
func makeBreakfast(item string, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("🍳 Preparing %s...\n", item)
	time.Sleep(duration) // Simulate preparation time
	fmt.Printf("🍽️ %s is ready!\n", item)
}

func main() {
	// ✅ Crucial for Parallelism: Utilize all available CPU cores
	numCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numCores)

	fmt.Printf("--- Parallel Breakfast Prep (Using %d CPU cores) ---\n", numCores)
	start := time.Now()
	var wg sync.WaitGroup

	// Let's make multiple breakfast items in parallel!
	wg.Add(1)
	go makeBreakfast("Eggs 🍳", 3*time.Second, &wg)

	wg.Add(1)
	go makeBreakfast("Pizza 🍕", 4*time.Second, &wg)

	wg.Add(1)
	go makeBreakfast("Pancakes 🥞", 2*time.Second, &wg)

	wg.Wait() // Wait for all breakfast items to be ready
	fmt.Printf("\n🎉 All 'parallel' breakfast items ready in %v!\n", time.Since(start))
	fmt.Println("Tasks ran simultaneously, making breakfast much faster!")
}
