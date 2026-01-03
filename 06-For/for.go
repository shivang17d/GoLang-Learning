package main

import "fmt"

func main() {
	// 1. Basic Loop (init; condition; post)
	// The standard for loop found in C/Java/JS
	fmt.Println("--- 1. Basic Loop ---")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// 2. Condition-Only Loop (While Loop)
	// Go has no 'while' keyword. Use 'for' with just a condition.
	fmt.Println("\n--- 2. While-style Loop ---")
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// 3. Infinite Loop
	// Use 'for' without any condition.
	fmt.Println("\n--- 3. Infinite Loop with Break ---")
	k := 0
	for {
		if k >= 3 {
			break // Exit the loop
		}
		fmt.Println(k)
		k++
	}

	// 4. Range Loop (Iteration)
	// Used for slices, arrays, maps, and strings.
	fmt.Println("\n--- 4. Range Loop (Slice) ---")
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 5. Range over Integer (Go 1.22+)
	// Iterates from 0 to n-1.
	fmt.Println("\n--- 5. Range over Integer ---")
	for i := range 3 { // Iterate 3 times (0, 1, 2)
		fmt.Println("Count:", i)
	}

	// 6. Break and Continue
	fmt.Println("\n--- 6. Break and Continue ---")
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue // Skip remainder of this iteration
		}
		if i == 3 {
			break // Exit loop entirely
		}
		fmt.Println(i)
	}
}
