package main

import "fmt"

func main() {
	fmt.Println("--- Range over Slices & Arrays ---")
	nums := []int{2, 3, 4}
	sum := 0
	// range on arrays and slices provides BOTH index and value
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
		sum += num
	}
	fmt.Println("Sum:", sum)

	// Use blank identifier '_' to ignore index
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println("Even number:", num)
		}
	}

	fmt.Println("\n--- Range over Maps ---")
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// range over maps provides key and value
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range over only keys
	fmt.Print("Keys only: ")
	for k := range kvs {
		fmt.Printf("%s ", k)
	}
	fmt.Println()

	fmt.Println("\n--- Range over Strings ---")
	// range on strings iterates over Unicode code points (runes).
	// The first value is the starting byte index of the rune.
	// The second value is the rune itself.
	for i, c := range "go🌏" {
		fmt.Printf("Byte Index: %d, Unicode: %c\n", i, c)
	}

	fmt.Println("\n--- Integer Range (Go 1.22+) ---")
	// You can now range over an integer directly.
	// It iterates from 0 to n-1.
	fmt.Print("Range 5: ")
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\n--- Value Semantics in Range ---")
	// IMPORTANT: The second value returned by range is a COPY of the element.
	// Modifying it does NOT change the original structure.
	arr := [3]int{1, 1, 1}
	for _, v := range arr {
		v = 10 // This only modifies the local copy 'v'
		_ = v  // Dummy use to satisfy compiler/lint
	}
	fmt.Println("Original after loop (unchanged):", arr)

	// To modify the original, use the index
	for i := range arr {
		arr[i] = 10
	}
	fmt.Println("Original after index loop (modified):", arr)
}
