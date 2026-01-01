package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We use bufio.NewReader for EVERYTHING to avoid buffer mixing issues.
	// fmt.Scan leaves newlines in the buffer, which is why your sentence was skipped.
	reader := bufio.NewReader(os.Stdin)

	// 1. Reading an Integer
	fmt.Print("Enter an integer: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer")
	} else {
		fmt.Println("You entered integer:", i)
	}

	// 2. Reading a Float
	fmt.Print("Enter a float: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error: Please enter a valid float")
	} else {
		fmt.Println("You entered float:", f)
	}

	// 3. Reading a Boolean
	fmt.Print("Enter a boolean (true/false): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	b, err := strconv.ParseBool(input)
	if err != nil {
		fmt.Println("Error: Please enter true or false")
	} else {
		fmt.Println("You entered boolean:", b)
	}

	// 4. Reading a String (Single word)
	fmt.Print("Enter a single word: ")
	input, _ = reader.ReadString('\n')
	// Even for single words, reading the full line and trimming is safer
	s := strings.Fields(input)[0]
	fmt.Println("You entered word:", s)

	// 5. Reading a full line stays exactly the same!
	fmt.Print("Enter a full sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)
	fmt.Println("You entered sentence:", sentence)
}
