package main

import "fmt"

func main() {
	fmt.Println("--- Array Declaration & Initialization ---")
	// 1. Declaration (Zero Values)
	// Declares an array of 5 integers. Default value is 0 for all elements.
	var arr1 [5]int
	fmt.Println("Arr1 (Integer Default):", arr1)

	var strArr [3]string
	fmt.Printf("String Array (Default): %q\n", strArr)

	var boolArr [3]bool
	fmt.Println("Boolean Array (Default):", boolArr)

	var floatArr [3]float64
	fmt.Println("Float Array (Default):", floatArr)

	// 2. Initialization with Literals
	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Arr2 (Literal):", arr2)

	// 3. Infer Length (...)
	// The compiler counts the elements to determine the length.
	arr3 := [...]int{1, 2, 3}
	fmt.Printf("Arr3 (Inferred Length): %v, Length: %d\n", arr3, len(arr3))

	// 4. Sparse Arrays (Specific Indices)
	// Initialize index 1 to 100, and index 4 to 200. Others are 0.
	arr4 := [5]int{1: 100, 4: 200}
	fmt.Println("Arr4 (Sparse):", arr4)

	fmt.Println("\n--- Accessing & Modifying ---")
	fmt.Println("Arr2[2] (Original):", arr2[2])
	arr2[2] = 300 // Modify
	fmt.Println("Arr2[2] (Modified):", arr2[2])
	fmt.Println("Arr2 (After Modify):", arr2)

	fmt.Println("\n--- Iterating ---")
	// Using standard for loop
	fmt.Print("For Loop: ")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%d ", arr2[i])
	}
	fmt.Println()

	// Using range (index and value)
	fmt.Print("Range (i, v): ")
	for i, v := range arr2 {
		fmt.Printf("[%d:%d] ", i, v)
	}
	fmt.Println()

	// Using range (value only)
	fmt.Print("Range (v): ")
	for _, v := range arr2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("\n--- Multi-Dimensional Arrays ---")
	// 2D Array: 2 rows, 3 columns
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Matrix[1][2]:", matrix[1][2]) // Access element at row 1, col 2

	fmt.Println("\n--- Value Semantics (Copying) ---")
	// Arrays in Go are values, not references.
	// Assigning one array to another COPIES all elements.
	original := [3]string{"A", "B", "C"}
	copyArr := original // Full copy happens here

	copyArr[0] = "Z" // Modifying copy does NOT affect original

	fmt.Println("Original:", original)
	fmt.Println("Copy:", copyArr)

	// Ensure types match strictly
	// var invalid [4]int = arr1 // This would be a COMPILE ERROR because [4]int and [5]int are different types.
}
