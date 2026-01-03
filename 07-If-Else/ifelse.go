package main

import "fmt"

func main() {
	// 1. Basic If-Else
	// No parentheses around the condition. Braces are mandatory.
	age := 18
	if age >= 18 {
		fmt.Println("1. Person is an adult.")
	} else {
		fmt.Println("1. Person is a minor.")
	}

	// 2. Else If Chain
	// Handling multiple conditions.
	score := 85
	if score >= 90 {
		fmt.Println("2. Grade: A")
	} else if score >= 80 {
		fmt.Println("2. Grade: B")
	} else {
		fmt.Println("2. Grade: C")
	}

	// 3. If with Short Statement
	// You can execute a statement before the condition.
	// Variables declared here are only available inside the if/else keys.
	if num := 9; num < 0 {
		fmt.Println("3. Number is negative:", num)
	} else if num < 10 {
		fmt.Println("3. Number is single digit:", num)
	} else {
		fmt.Println("3. Number has multiple digits:", num)
	}
	// fmt.Println(num) // Error: num is undefined here

	// 4. Logical Operators
	// && (AND), || (OR), ! (NOT)
	role := "admin"
	hasPermission := true

	if role == "admin" || hasPermission {
		fmt.Println("4. Access Granted")
	}

	if role == "admin" && hasPermission {
		fmt.Println("4. Super User Access Granted")
	}

	if !hasPermission {
		fmt.Println("4. Access Denied")
	}
}
