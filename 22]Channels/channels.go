package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- Channels ---")

	// 1. Unbuffered Channel (block until both sender and receiver are ready)
	messages := make(chan string)

	go func() {
		messages <- "ping" // Send to channel
	}()

	msg := <-messages // Receive from channel
	fmt.Println("Received from unbuffered channel:", msg)

	// 2. Buffered Channel (doesn't block sender until buffer is full)
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println("Sent 1 and 2 to buffered channel")
	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)

	// 3. Channel Closing and Range
	jobs := make(chan int, 5)
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs) // Closing tells receivers no more values will be sent

	fmt.Println("Reading closed channel using range:")
	for j := range jobs {
		fmt.Println("Job:", j)
	}

	// 4. Select Statement (Multiplexing)
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	fmt.Println("Waiting for messages using select...")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from c2:", msg2)
		}
	}
}
