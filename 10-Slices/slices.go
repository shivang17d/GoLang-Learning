package main

import (
	"fmt"
	"slices"
)

func main() {
	// IMPORTANT: Slices are Go's version of a "Dynamic Array".
	// Unlike Arrays, Slices can grow and shrink in size.
	fmt.Println("--- Slices: Go's Dynamic Arrays ---")
	// 1. Literal Pattern
	// Slices do NOT have a fixed length in the type (e.g., []int vs [5]int)
	slice1 := []int{1, 2, 3}
	fmt.Printf("Slice1 (Literal): %v, Len: %d, Cap: %d\n", slice1, len(slice1), cap(slice1))

	// 2. Using make()
	// make([]Type, len, cap). Cap is optional (defaults to len). Can resize.
	// Creates an underlying array of size 5, returns a slice view of size 3.
	slice2 := make([]int, 3, 5)
	slice2[0] = 10
	slice2[1] = 20
	slice2[2] = 30
	fmt.Printf("Slice2 (make): %v, Len: %d, Cap: %d\n", slice2, len(slice2), cap(slice2))

	// 3. Slicing an existing Array (or Slice)
	arr := [5]int{100, 200, 300, 400, 500}
	slice3 := arr[1:4] // Index 1 (inclusive) to 4 (exclusive) -> [200, 300, 400]
	fmt.Printf("Slice3 (from Array): %v, Len: %d, Cap: %d\n", slice3, len(slice3), cap(slice3))

	// 4. Three-index Slicing (Go 1.2+)
	// s[low:high:max] - 'max' controls the capacity of the resulting slice
	slice5 := arr[1:2:3] // Starts at 1, length is (2-1)=1, capacity is (3-1)=2
	fmt.Printf("Slice5 (3-index): %v, Len: %d, Cap: %d\n", slice5, len(slice5), cap(slice5))

	fmt.Println("\n--- Reference Semantics (Slice vs Array) ---")
	// Slices are "windows" to an underlying array.
	fmt.Println("Original Array:", arr)
	slice3[0] = 999
	fmt.Println("Modified Slice3:", slice3)
	fmt.Println("Array After Slice Mod:", arr) // Array IS modified!

	fmt.Println("\n--- Dynamic Growth (Append) ---")
	// append() returns a NEW slice header.
	// If capacity is sufficient, it uses the same underlying array.
	// If not, it allocates a new larger array and copies data.
	var slice4 []int
	fmt.Printf("Initial: %v, Len: %d, Cap: %d\n", slice4, len(slice4), cap(slice4))

	for i := 0; i < 5; i++ {
		slice4 = append(slice4, i)
		fmt.Printf("Append %d: %v, Len: %d, Cap: %d\n", i, slice4, len(slice4), cap(slice4))
	}

	// 2. Variadic Append (Merging two slices)
	// The ... operator "unpacks" the slice into individual elements
	otherSlice := []int{100, 200, 300}
	slice4 = append(slice4, otherSlice...)
	fmt.Println("After Merging otherSlice:", slice4)

	fmt.Println("\n--- Iterating over Slices ---")
	// Using range is the standard way to iterate
	for i, v := range slice4 {
		fmt.Printf("Index: %d, Value: %d | ", i, v)
		if i == 2 {
			fmt.Println("...")
			break
		}
	}
	fmt.Println()

	fmt.Println("\n--- Copying Slices ---")
	// copy(dst, src) copies min(len(dst), len(src)) elements.
	// Use this to create an INDEPENDENT copy.
	src := []string{"A", "B", "C"}
	dst := make([]string, len(src))
	count := copy(dst, src)
	fmt.Printf("Copied %d elements. Dst: %v\n", count, dst)

	dst[0] = "Z"
	fmt.Println("Modified Dst:", dst)
	fmt.Println("Original Src:", src) // UNCHANGED

	fmt.Println("\n--- Nil vs Empty Slice ---")
	var nilSlice []int    // Zero value is nil
	emptySlice := []int{} // Initialized but empty, non-nil

	fmt.Printf("Nil Slice: %v, Is Nil? %t, Len: %d\n", nilSlice, nilSlice == nil, len(nilSlice))
	fmt.Printf("Empty Slice: %v, Is Nil? %t, Len: %d\n", emptySlice, emptySlice == nil, len(emptySlice))

	fmt.Println("\n--- Built-in Function: clear() (Go 1.21+) ---")
	// clear() zeros out all elements but keeps the length/capacity
	clearTarget := []int{10, 20, 30, 40}
	fmt.Println("Before clear:", clearTarget)
	clear(clearTarget)
	fmt.Println("After clear: ", clearTarget, "Len:", len(clearTarget))

	fmt.Println("\n--- 'slices' Package Functions (Go 1.21+) ---")
	s1 := []int{50, 20, 90, 10, 40}

	// 1. Sort
	slices.Sort(s1)
	fmt.Println("Sorted:", s1)

	// 2. Contains & Index
	fmt.Println("Contains 20?", slices.Contains(s1, 20))
	fmt.Println("Index of 90:", slices.Index(s1, 90))

	// 3. Reverse
	slices.Reverse(s1)
	fmt.Println("Reversed:", s1)

	// 4. Delete (removes elements at index/range)
	// Delete elements from index 1 to 3 (exclusive)
	s1 = slices.Delete(s1, 1, 3)
	fmt.Println("After Delete [1:3]:", s1)

	// 5. Insert
	s1 = slices.Insert(s1, 1, 999, 888)
	fmt.Println("After Insert 999, 888 at index 1:", s1)

	// 6. Max & Min
	fmt.Println("Max:", slices.Max(s1))
	fmt.Println("Min:", slices.Min(s1))

	fmt.Println("\n--- Classic Idioms (Using only Built-ins) ---")
	// Use these if you cannot use the 'slices' package or need custom logic

	// 1. Delete element at index i (e.g., index 1)
	// s = append(s[:i], s[i+1:]...)
	nums := []int{1, 2, 3, 4, 5}
	i := 1
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println("Classic Delete index 1:", nums)

	// 2. Filter (In-place) - Keep even numbers
	// This is very memory efficient as it uses the same underlying array
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := 0
	for _, x := range nums2 {
		if x%2 == 0 {
			nums2[n] = x
			n++
		}
	}
	nums2 = nums2[:n]
	fmt.Println("In-place Filter (Even):", nums2)

	fmt.Println("\n--- 2D Slices (Multi-dimensional) ---")
	// 1. Literal Pattern
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D Slice (Literal):", matrix)

	// 2. Dynamic Initialization (Jagged Slices)
	// Slices of slices can have different lengths for each inner slice.
	jagged := make([][]int, 3)
	for i := 0; i < len(jagged); i++ {
		innerLen := i + 1
		jagged[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			jagged[i][j] = i + j
		}
	}
	fmt.Println("Jagged 2D Slice:")
	for _, row := range jagged {
		fmt.Println("  ", row)
	}
}
