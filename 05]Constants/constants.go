package main

import "fmt"

// Package-level constants and variables
const pi = 3.14         // Untyped constant (available throughout the package)
var globalName = "User" // Variable (can be changed)

func main() {
	fmt.Println("--- 1. Package Level Constants ---")
	fmt.Println("Value of pi:", pi)
	fmt.Println("Global Name:", globalName)

	fmt.Println("\n--- 2. Constants vs Variables ---")
	// Constants must be initialized at declaration and cannot change.
	const name = "Golang" // This 'name' shadows any global variables if they had the same name
	fmt.Println("Constant Name:", name)

	// name = "Python" // Error: cannot assign to name (constants are immutable)

	// Variables can be reassigned.
	globalName = "Power User"
	fmt.Println("Updated Global Name:", globalName)

	fmt.Println("\n--- 3. Typed vs Untyped Constants ---")
	const typedInt int = 10 // Typed constant (strictly int)
	const untypedInt = 10   // Untyped constant (interpreted as needed)

	fmt.Println("Typed:", typedInt)
	fmt.Println("Untyped:", untypedInt)

	// Untyped constants allow for flexibility
	var myFloat float64 = untypedInt // Works because untypedInt is just a numeric value "10"
	fmt.Println("Untyped assigned to float:", myFloat)

	// var myFloat2 float64 = typedInt // Error: cannot use typedInt (type int) as type float64 in assignment

	fmt.Println("\n--- 4. Multiple Declaration ---")
	const (
		port = 8080
		host = "localhost"
	)
	fmt.Printf("Server running at %s:%d\n", host, port)

	fmt.Println("\n--- 5. Iota (Enumerations) ---")
	// iota allows creating sequential constants
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println("Days of week:", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println("\n--- 6. Iota Expressions ---")
	// Using iota in expressions (e.g., bitwise shifting for storage sizes)
	const (
		_  = iota             // Skip 0 (0 * 10 = 0)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30
	)
	fmt.Printf("KB: %d, MB: %d, GB: %d\n", KB, MB, GB)
}
