package main

import "fmt"

func main() {
	fmt.Println("--- Basic State Capture ---")
	// Closures capture the environment in which they are created.
	// Here, 'x' persists between calls to the returned function.
	counter := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
	}

	c1 := counter()
	fmt.Println("c1:", c1()) // 1
	fmt.Println("c1:", c1()) // 2
	fmt.Println("c1:", c1()) // 3

	fmt.Println("\n--- Isolated State (Multiple Instances) ---")
	// Each call to the factory function creates a NEW isolated scope.
	c2 := counter()
	fmt.Println("c2:", c2())       // 1 (new state)
	fmt.Println("c1 again:", c1()) // 4 (old state preserved)

	fmt.Println("\n--- Higher-Order Functions (Passing Closures) ---")
	// Using a closure as a custom "predicate" or "filter".
	nums := []int{1, 2, 3, 4, 5, 6}

	filter := func(arr []int, test func(int) bool) []int {
		result := []int{}
		for _, v := range arr {
			if test(v) {
				result = append(result, v)
			}
		}
		return result
	}

	// Define 'isEven' as a closure
	evens := filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	fmt.Println("\n--- Closures in Loops (Go 1.22+ Behavior) ---")
	// IMPORTANT: Pre-Go 1.22, 'i' was shared across iterations, often causing bugs.
	// In Go 1.22+, each iteration has its OWN 'i', making closure capture safe.
	funcs := make([]func(), 3)
	for i := range 3 {
		funcs[i] = func() {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Print("Loop Execution (Expected 0 1 2): ")
	for _, f := range funcs {
		f()
	}
	fmt.Println()

	fmt.Println("\n--- Wrapping State ---")
	// A closure can "wrap" logic and state together.
	adder := func(base int) func(int) int {
		return func(val int) int {
			return base + val
		}
	}

	add10 := adder(10)
	add20 := adder(20)

	fmt.Println("10 + 5 =", add10(5))
	fmt.Println("20 + 5 =", add20(5))
}
