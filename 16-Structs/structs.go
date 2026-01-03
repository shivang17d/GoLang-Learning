package main

import "fmt"

// 1. Basic Struct Definition
type ContactInfo struct {
	Email   string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName  string
	// 2. Nested Struct
	Contact ContactInfo
}

// 3. Struct Embedding (Composition)
// This is NOT inheritance, but it looks like it.
// Fields of ContactInfo are "promoted" to Student.
type Student struct {
	ID int
	ContactInfo
}

func main() {
	fmt.Println("--- Struct Initialization ---")

	// Way 1: Zero Value (fields initialized to their zero values)
	var p1 Person
	fmt.Printf("p1 (Zero Value): %+v\n", p1)

	// Way 2: Struct Literal (by order, not recommended)
	p2 := Person{"Alice", "Smith", ContactInfo{"alice@example.com", 12345}}
	fmt.Printf("p2 (Literal): %+v\n", p2)

	// Way 3: Struct Literal (with field names - BEST PRACTICE)
	p3 := Person{
		FirstName: "Bob",
		LastName:  "Jones",
		Contact: ContactInfo{
			Email:   "bob@example.com",
			ZipCode: 54321,
		},
	}
	fmt.Printf("p3 (Field Names): %+v\n", p3)

	// Way 4: Pointers to Structs
	p4 := &Person{FirstName: "Charlie"}
	fmt.Printf("p4 (Pointer): %p, Value: %+v\n", p4, *p4)

	fmt.Println("\n--- Accessing & Modifying Fields ---")
	p3.FirstName = "Robert"
	fmt.Println("p3 New Name:", p3.FirstName)
	fmt.Println("p3 Zip Code:", p3.Contact.ZipCode)

	fmt.Println("\n--- Anonymous Structs ---")
	// Used for one-off data structures, like JSON responses in tests
	config := struct {
		APIKey string
		Port   int
	}{
		APIKey: "secret-key",
		Port:   8080,
	}
	fmt.Printf("Anonymous Config: %+v\n", config)

	fmt.Println("\n--- Struct Embedding (Composition) ---")
	s1 := Student{
		ID: 101,
		ContactInfo: ContactInfo{
			Email:   "student@school.edu",
			ZipCode: 90210,
		},
	}

	// Promotion: We can access ContactInfo fields directly on s1!
	fmt.Println("Student ID:", s1.ID)
	fmt.Println("Student Email (Promoted):", s1.Email)
	fmt.Println("Student Zip (Promoted):", s1.ZipCode)
	// We can still access the inner struct explicitly
	fmt.Println("Inner Contact Email:", s1.ContactInfo.Email)
}
