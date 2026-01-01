package utils

import "fmt"

// Exported function (starts with Uppercase)
func SayHello(name string) {
	fmt.Printf("Hello, %s! (from utils package)\n", name)
}

// Unexported function (starts with lowercase)
func internalHelper() {
	fmt.Println("This is an internal helper.")
}

// Exported function calling unexported one
func UseHelper() {
	internalHelper()
}
