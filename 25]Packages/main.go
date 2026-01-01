package main

import (
	"fmt"
	"myproject/utils" // Importing the local package
)

func main() {
	fmt.Println("--- Multi-Package Project ---")

	utils.SayHello("Go Developer")

	// utils.internalHelper() // This would fail (unexported)

	utils.UseHelper() // This works
}
