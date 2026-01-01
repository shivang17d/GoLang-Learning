package main

import "fmt"

// 1. Basic Function
// plus accepts two integers and returns an integer result.
func plus(a int, b int) int {
	return a + b
}

// 2. Multiple Return Values
// manyReturns returns two results.
func manyReturns() (int, int) {
	return 3, 7
}

// 3. Named Return Values
// result is implicitly returned.
func namedReturn(a, b int) (result int) {
	result = a * b
	return // Naked return
}

// 4. Variadic Functions
// sum accepts any number of integers.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 5. Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("--- Basic Functions ---")
	res := plus(5, 10)
	fmt.Println("5 + 10 =", res)

	fmt.Println("\n--- Multiple Return Values ---")
	a, b := manyReturns()
	fmt.Printf("Returned: %d and %d\n", a, b)

	// Ignore one of the values using _
	_, c := manyReturns()
	fmt.Println("Second value only:", c)

	fmt.Println("\n--- Named Return Values ---")
	fmt.Println("3 * 4 =", namedReturn(3, 4))

	fmt.Println("\n--- Variadic Functions ---")
	fmt.Println("Sum(1, 2):", sum(1, 2))
	fmt.Println("Sum(1, 2, 3, 4):", sum(1, 2, 3, 4))

	// Appending a slice to a variadic function
	nums := []int{10, 20, 30}
	fmt.Println("Sum(slice):", sum(nums...))

	fmt.Println("\n--- Anonymous Functions & Closures ---")
	// Anonymous function assigned to a variable
	addOne := func(i int) int {
		return i + 1
	}
	fmt.Println("addOne(5):", addOne(5))

	// Closures: Function that returns another function
	// and "captures" the local variable 'i'.
	intSeq := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	nextInt := intSeq()
	fmt.Print("Sequence: ")
	fmt.Printf("%d, ", nextInt())
	fmt.Printf("%d, ", nextInt())
	fmt.Printf("%d\n", nextInt())

	fmt.Println("\n--- Recursion ---")
	fmt.Println("7! =", factorial(7))
}
