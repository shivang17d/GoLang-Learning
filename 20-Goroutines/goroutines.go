package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("--- Goroutines ---")

	// 1. Basic Goroutine
	// Running this function in the background.
	go say("world")

	// Running this one in the foreground (main goroutine).
	say("hello")

	// 2. Anonymous Goroutine
	go func(msg string) {
		fmt.Println("Anonymous says:", msg)
	}("I'm running in the background!")

	// Give background goroutines a moment to finish before main exits.
	// In a real app, we use WaitGroups or Channels for this!
	time.Sleep(1 * time.Second)
	fmt.Println("Main finished.")
}
