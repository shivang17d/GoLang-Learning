package main

import (
	"fmt"
	"maps" // Go 1.21+
)

func main() {
	fmt.Println("--- Map Creation ---")
	// 1. Literal Pattern
	// map[KeyType]ValueType
	m1 := map[string]int{
		"apple":  10,
		"banana": 20,
	}
	fmt.Println("Map1 (Literal):", m1)

	// 2. Using make()
	// make(map[KeyType]ValueType, initialCapacity)
	// initialCapacity is optional but helps with performance
	m2 := make(map[string]string)
	m2["name"] = "Alice"
	m2["city"] = "Wonderland"
	fmt.Println("Map2 (make):", m2)

	fmt.Println("\n--- Accessing & Updating ---")
	// Update
	m1["apple"] = 15
	// Add
	m1["orange"] = 30
	fmt.Println("Map1 updated:", m1)

	// Accessing a non-existent key returns the zero value of the ValueType (integer)
	fmt.Println("Non-existent key (integer):", m1["pear"]) // Output: 0

	fmt.Println("\n--- Existence Check (ok idiom) ---")
	// val, ok := map[key]
	// 'ok' is true if the key exists, false otherwise.
	val, ok := m1["banana"]
	if ok {
		fmt.Printf("Banana found: %d\n", val)
	} else {
		fmt.Println("Banana not found")
	}

	fmt.Println("\n--- Deleting ---")
	// delete(map, key) - No error if key doesn't exist.
	delete(m1, "orange")
	fmt.Println("Map1 after delete 'orange':", m1)

	fmt.Println("\n--- Iterating ---")
	// Map iteration order is RANDOM (not guaranteed).
	m3 := map[int]string{1: "One", 2: "Two", 3: "Three"}
	for k, v := range m3 {
		fmt.Printf("[%d:%s] ", k, v)
	}
	fmt.Println()

	fmt.Println("\n--- Reference Semantics ---")
	// Maps are reference types. Modifying a copy affects the original.
	original := map[string]int{"A": 1}
	copyMap := original
	copyMap["B"] = 2
	fmt.Println("Original:", original)
	fmt.Println("Copy (modified):", copyMap)

	fmt.Println("\n--- Modern Features (Go 1.21+) ---")
	// 1. clear() - Removes all elements, leaves empty map
	fmt.Print("Before clear: ", m1)
	clear(m1)
	fmt.Println(" | After clear:", m1)

	// 2. maps package
	m4 := map[string]int{"X": 100, "Y": 200}
	m5 := map[string]int{"X": 100, "Y": 200}

	// Equal check
	fmt.Println("m4 == m5?", maps.Equal(m4, m5))

	// Clone
	m6 := maps.Clone(m4)
	fmt.Println("m6 (cloned from m4):", m6)

	fmt.Println("\n--- Idiom: Maps as Sets ---")
	// Go doesn't have a built-in Set type.
	// The common idiom is using a map[Type]struct{} or map[Type]bool.
	// struct{} is preferred as it takes zero extra memory.
	set := make(map[string]struct{})

	// Add to set
	set["user1"] = struct{}{}
	set["user2"] = struct{}{}

	// Check if in set
	_, exists := set["user1"]
	fmt.Println("user1 exists in set?", exists)

	_, exists3 := set["user3"]
	fmt.Println("user3 exists in set?", exists3)
}
