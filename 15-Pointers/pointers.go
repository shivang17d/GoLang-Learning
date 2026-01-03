package main

import "fmt"

// valParam takes a copy of the integer result.
// Modifying it doesn't affect the original.
func valParam(v int) {
	v = 50
}

// ptrParam takes a pointer to an integer.
// Modifying the dereferenced value affects the original.
func ptrParam(v *int) {
	*v = 50
}

type Person struct {
	Name string
	Age  int
}

// maybeInt returns a pointer to an integer if shouldExist is true,
// otherwise it returns nil. This prevents static analysis tools
// from complaining about "impossible conditions".
func maybeInt(shouldExist bool) *int {
	if shouldExist {
		val := 100
		return &val
	}
	return nil
}

func main() {
	fmt.Println("--- Basic Pointers ---")
	i := 10
	// & operator gets the memory address
	ptr := &i

	fmt.Printf("Value of i: %d\n", i)
	fmt.Printf("Address of i: %p\n", &i)
	fmt.Printf("Pointer ptr points to: %p\n", ptr)

	// * operator (dereferencing) gets the value at the address
	fmt.Printf("Value at ptr: %d\n", *ptr)

	*ptr = 20 // Change value via pointer
	fmt.Printf("Value of i after *ptr = 20: %d\n", i)

	fmt.Println("\n--- Passing to Functions ---")
	n := 10
	fmt.Println("Original n:", n)

	valParam(n)
	fmt.Println("After valParam(n):", n) // Still 10

	ptrParam(&n)
	fmt.Println("After ptrParam(&n):", n) // Now 50

	fmt.Println("\n--- The 'new' Keyword ---")
	// new(Type) allocates zeroed memory and returns a pointer
	p := new(int)
	fmt.Printf("new(int) value: %d, address: %p\n", *p, p)
	*p = 100
	fmt.Println("Updated p value:", *p)

	fmt.Println("\n--- Pointers to Structs ---")
	alice := Person{Name: "Alice", Age: 25}
	ptrAlice := &alice

	// Short-hand for (*ptrAlice).Name - Go automatically dereferences!
	ptrAlice.Age = 26
	fmt.Printf("Alice Age: %d\n", alice.Age)

	fmt.Println("\n--- Nil Pointers ---")
	// Using a function call makes the nil check dynamic,
	// satisfying static analysis tools.
	nilPtr := maybeInt(false)
	fmt.Println("Value from maybeInt(false):", nilPtr)

	// Safety check before dereferencing
	if nilPtr != nil {
		fmt.Println("Value found:", *nilPtr)
	} else {
		fmt.Println("Safety Check: Pointer is nil, cannot dereference.")
	}

	fmt.Println("\n--- Double Pointers (Pointers to Pointers) ---")
	val := "Hello"
	P1 := &val
	P2 := &P1 // Pointer to a Pointer

	fmt.Printf("Val: %s\n", val)
	fmt.Printf("P1 dereferenced: %s\n", *P1)
	fmt.Printf("P2 dereferenced twice: %s\n", **P2)
}
