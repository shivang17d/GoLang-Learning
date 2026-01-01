package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("--- File Handling ---")

	content := "Hello, Go File Handling!\nLine 2\nLine 3"
	filename := "test.txt"

	// 1. Writing a whole file
	err := os.WriteFile(filename, []byte(content), 0644)
	check(err)
	fmt.Println("Wrote to", filename)

	// 2. Reading a whole file
	data, err := os.ReadFile(filename)
	check(err)
	fmt.Println("Read whole file content:", string(data))

	// 3. Streaming large files (using Open and Scanner)
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	fmt.Println("Reading line by line:")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
	check(scanner.Err())

	// 4. Appending to a file
	f_append, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	_, err = f_append.WriteString("Appended Line\n")
	check(err)
	f_append.Close()
	fmt.Println("Appended a line.")

	// 5. Cleanup (Deleting the file)
	// err = os.Remove(filename)
	// check(err)
	// fmt.Println("Deleted", filename)
}
