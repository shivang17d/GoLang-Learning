package main

import "fmt"

func main() {
	// INTEGER
	// Integers are whole numbers.
	// We can simply print the integer.
	fmt.Println(1)
	// We can perform arithmetic operations.
	fmt.Println(1 + 1)

	// STRING
	// Strings are text sequences enclosed in double quotes.
	fmt.Println("Hello Golang")
	// We can concatenate strings using the + operator.
	fmt.Println("Hello " + "World")

	// Note : Single quotes doesn't work in go for strings.

	// BOOLEAN
	// Booleans represent true or false values.
	fmt.Println(true)
	fmt.Println(false)

	// FLOAT
	// Floats are numbers with decimal points.
	fmt.Println(10.5)
	// Arithmetic with floats works as expected.
	fmt.Println(7.0 / 3.0)

}
