package main

import "fmt"

// 1. Generic Function
// T is a type parameter. [T any] means T can be any type.
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// 2. Type Constraints
// 'comparable' is a built-in constraint for types that support == and !=
func FindIndex[T comparable](s []T, val T) int {
	for i, v := range s {
		if v == val {
			return i
		}
	}
	return -1
}

// 3. Custom Type Constraints
// We can define an interface that restricts the types allowed.
type Number interface {
	int | int64 | float64
}

func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// 4. Generic Structs
type Box[T any] struct {
	Content T
}

func main() {
	fmt.Println("--- Generic Functions ---")
	ints := []int{1, 2, 3}
	strs := []string{"Go", "Generics", "Fun"}

	PrintSlice(ints) // Type inference: Go knows T is int
	PrintSlice(strs) // Type inference: Go knows T is string

	fmt.Println("\n--- Using 'comparable' ---")
	fmt.Println("Index of 2:", FindIndex(ints, 2))
	fmt.Println("Index of 'Generics':", FindIndex(strs, "Generics"))

	fmt.Println("\n--- Custom Constraints (Number) ---")
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println("Sum of ints:", Sum(ints))
	fmt.Println("Sum of floats:", Sum(floats))

	fmt.Println("\n--- Generic Structs ---")
	intBox := Box[int]{Content: 100}
	strBox := Box[string]{Content: "Secret Message"}

	fmt.Printf("Int Box: %+v\n", intBox)
	fmt.Printf("String Box: %+v\n", strBox)
}
