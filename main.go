package main

import (
	"fmt"
	"runtime"
	"time"
)

func performTask(taskName string, duration time.Duration) {
	fmt.Printf("➡️ Solo Chef starting %s...\n", taskName)
	time.Sleep(duration) // Simulating work
	fmt.Printf("✅ Solo Chef finished %s!\n", taskName)
}

func main() {
	// ⚠️ Restricting to 1 CPU core to force task-switching (Concurrency)
	runtime.GOMAXPROCS(1)

	fmt.Println("--- The Concurrent Morning Routine (Solo Chef) ---")
	start := time.Now()

	go performTask("Toast 🍞", 2*time.Second)
	go performTask("Eggs 🍳", 3*time.Second)
	go performTask("Coffee ☕", 1*time.Second)

	// Wait for background tasks to finish
	time.Sleep(4 * time.Second)
	fmt.Printf("\n✨ Tasks managed in %v!\n", time.Since(start))
	fmt.Println("One person interleaved tasks to get everything done.")
}
