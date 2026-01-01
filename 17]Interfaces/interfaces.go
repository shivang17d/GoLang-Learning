package main

import (
	"fmt"
	"math"
)

// 1. Interface Declaration
// An interface defines a set of method signatures.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Implementation
// In Go, interfaces are implemented IMPLICITLY.
// There is no "implements" keyword. A type implements an interface
// by simply defining all the methods in that interface.

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. Polymorphism
// This function takes ANY type that satisfies the Shape interface.
func printShapeInfo(s Shape) {
	fmt.Printf("Type: %T, Area: %.2f, Perimeter: %.2f\n", s, s.Area(), s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	fmt.Println("--- Basic Interfaces ---")
	printShapeInfo(c)
	printShapeInfo(r)

	fmt.Println("\n--- Empty Interface (interface{} or any) ---")
	// An empty interface can hold any value (since all types have zero or more methods)
	var anything interface{} // In modern Go, you can also use `any`
	anything = "Hello"
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)
	anything = 42
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	fmt.Println("\n--- Type Assertions ---")
	// Used to get the underlying concrete value from an interface
	var s Shape = Circle{Radius: 10}

	// Safe type assertion
	if c2, ok := s.(Circle); ok {
		fmt.Printf("Success! Radius is %.2f\n", c2.Radius)
	} else {
		fmt.Println("Assertion failed: Not a Circle")
	}

	fmt.Println("\n--- Type Switches ---")
	// A cleaner way to check multiple types
	describeValue(10)
	describeValue("Go is fun")
	describeValue(c)
}

func describeValue(v any) {
	switch val := v.(type) {
	case int:
		fmt.Printf("It's an integer: %d\n", val)
	case string:
		fmt.Printf("It's a string: %s\n", val)
	case Shape:
		fmt.Printf("It's a Shape with area: %.2f\n", val.Area())
	default:
		fmt.Printf("Unknown type: %T\n", val)
	}
}
